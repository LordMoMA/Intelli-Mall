package baskets

import (
	"context"
	"database/sql"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rs/zerolog"

	"github.com/LordMoMA/Intelli-Mall/baskets/basketspb"
	"github.com/LordMoMA/Intelli-Mall/baskets/internal/application"
	"github.com/LordMoMA/Intelli-Mall/baskets/internal/constants"
	"github.com/LordMoMA/Intelli-Mall/baskets/internal/domain"
	"github.com/LordMoMA/Intelli-Mall/baskets/internal/grpc"
	"github.com/LordMoMA/Intelli-Mall/baskets/internal/handlers"
	"github.com/LordMoMA/Intelli-Mall/baskets/internal/postgres"
	"github.com/LordMoMA/Intelli-Mall/baskets/internal/rest"
	"github.com/LordMoMA/Intelli-Mall/internal/am"
	"github.com/LordMoMA/Intelli-Mall/internal/amotel"
	"github.com/LordMoMA/Intelli-Mall/internal/amprom"
	"github.com/LordMoMA/Intelli-Mall/internal/ddd"
	"github.com/LordMoMA/Intelli-Mall/internal/di"
	"github.com/LordMoMA/Intelli-Mall/internal/es"
	"github.com/LordMoMA/Intelli-Mall/internal/jetstream"
	pg "github.com/LordMoMA/Intelli-Mall/internal/postgres"
	"github.com/LordMoMA/Intelli-Mall/internal/postgresotel"
	"github.com/LordMoMA/Intelli-Mall/internal/registry"
	"github.com/LordMoMA/Intelli-Mall/internal/system"
	"github.com/LordMoMA/Intelli-Mall/internal/tm"
	"github.com/LordMoMA/Intelli-Mall/stores/storespb"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono system.Service) (err error) {
	return Root(ctx, mono)
}

func Root(ctx context.Context, svc system.Service) (err error) {
	// Container Setup: Create a dependency injection container.
	container := di.New()
	// setup Driven adapters
	container.AddSingleton(constants.RegistryKey, func(c di.Container) (any, error) {
		reg := registry.New()
		if err := domain.Registrations(reg); err != nil {
			return nil, err
		}
		if err := basketspb.Registrations(reg); err != nil {
			return nil, err
		}
		if err := storespb.Registrations(reg); err != nil {
			return nil, err
		}
		return reg, nil
	})
	stream := jetstream.NewStream(svc.Config().Nats.Stream, svc.JS(), svc.Logger())
	container.AddSingleton(constants.DomainDispatcherKey, func(c di.Container) (any, error) {
		return ddd.NewEventDispatcher[ddd.Event](), nil
	})
	container.AddScoped(constants.DatabaseTransactionKey, func(c di.Container) (any, error) {
		return svc.DB().Begin()
	})
	sentCounter := amprom.SentMessagesCounter(constants.ServiceName)
	container.AddScoped(constants.MessagePublisherKey, func(c di.Container) (any, error) {
		tx := postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx))
		outboxStore := pg.NewOutboxStore(constants.OutboxTableName, tx)
		return am.NewMessagePublisher(
			stream,
			amotel.OtelMessageContextInjector(),
			sentCounter,
			tm.OutboxPublisher(outboxStore),
		), nil
	})
	container.AddSingleton(constants.MessageSubscriberKey, func(c di.Container) (any, error) {
		return am.NewMessageSubscriber(
			stream,
			amotel.OtelMessageContextExtractor(),
			amprom.ReceivedMessagesCounter(constants.ServiceName),
		), nil
	})
	container.AddScoped(constants.EventPublisherKey, func(c di.Container) (any, error) {
		return am.NewEventPublisher(
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.MessagePublisherKey).(am.MessagePublisher),
		), nil
	})
	container.AddScoped(constants.InboxStoreKey, func(c di.Container) (any, error) {
		tx := postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx))
		return pg.NewInboxStore(constants.InboxTableName, tx), nil
	})
	container.AddScoped(constants.BasketsRepoKey, func(c di.Container) (any, error) {
		tx := postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx))
		reg := c.Get(constants.RegistryKey).(registry.Registry)
		return es.NewAggregateRepository[*domain.Basket](
			domain.BasketAggregate,
			reg,
			es.AggregateStoreWithMiddleware(
				pg.NewEventStore(constants.EventsTableName, tx, reg),
				pg.NewSnapshotStore(constants.SnapshotsTableName, tx, reg),
			),
		), nil
	})
	container.AddScoped(constants.StoresRepoKey, func(c di.Container) (any, error) {
		return postgres.NewStoreCacheRepository(
			constants.StoresCacheTableName,
			postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx)),
			grpc.NewStoreRepository(svc.Config().Rpc.Service(constants.StoresServiceName)),
		), nil
	})
	container.AddScoped(constants.ProductsRepoKey, func(c di.Container) (any, error) {
		return postgres.NewProductCacheRepository(
			constants.ProductsCacheTableName,
			postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx)),
			grpc.NewProductRepository(svc.Config().Rpc.Service(constants.StoresServiceName)),
		), nil
	})
	// Prometheus counters
	var basketsStarted = promauto.NewCounter(prometheus.CounterOpts{
		Name: constants.BasketsStartedCount,
	})
	basketsCheckedOut := promauto.NewCounter(prometheus.CounterOpts{
		Name: constants.BasketsCheckedOutCount,
	})
	basketsCanceled := promauto.NewCounter(prometheus.CounterOpts{
		Name: constants.BaksetsCanceledCount,
	})

	// setup application
	// Injecting dependencies into the application.
	container.AddScoped(constants.ApplicationKey, func(c di.Container) (any, error) {
		return application.NewInstrumentedApp(application.New(
			c.Get(constants.BasketsRepoKey).(domain.BasketRepository),              // Inject BasketRepository.
			c.Get(constants.StoresRepoKey).(domain.StoreCacheRepository),           // Inject StoreCacheRepository.
			c.Get(constants.ProductsRepoKey).(domain.ProductCacheRepository),       // Inject ProductCacheRepository.
			c.Get(constants.DomainDispatcherKey).(*ddd.EventDispatcher[ddd.Event]), // Inject EventDispatcher.
		), basketsStarted, basketsCheckedOut, basketsCanceled), nil
	})
	// Injecting dependencies into event handlers.
	container.AddScoped(constants.DomainEventHandlersKey, func(c di.Container) (any, error) {
		return handlers.NewDomainEventHandlers(c.Get(constants.EventPublisherKey).(am.EventPublisher)), nil
	})
	container.AddScoped(constants.IntegrationEventHandlersKey, func(c di.Container) (any, error) {
		return handlers.NewIntegrationEventHandlers(
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.StoresRepoKey).(domain.StoreCacheRepository),
			c.Get(constants.ProductsRepoKey).(domain.ProductCacheRepository),
			tm.InboxHandler(c.Get(constants.InboxStoreKey).(tm.InboxStore)),
		), nil
	})
	outboxProcessor := tm.NewOutboxProcessor(
		stream,
		pg.NewOutboxStore(constants.OutboxTableName, svc.DB()),
	)

	// setup Driver adapters
	if err = grpc.RegisterServerTx(container, svc.RPC()); err != nil {
		return err
	}
	if err = rest.RegisterGateway(ctx, svc.Mux(), svc.Config().Rpc.Address()); err != nil {
		return err
	}
	if err = rest.RegisterSwagger(svc.Mux()); err != nil {
		return err
	}
	handlers.RegisterDomainEventHandlersTx(container)
	if err = handlers.RegisterIntegrationEventHandlersTx(container); err != nil {
		return err
	}
	startOutboxProcessor(ctx, outboxProcessor, svc.Logger())
	return
}

func startOutboxProcessor(ctx context.Context, outboxProcessor tm.OutboxProcessor, logger zerolog.Logger) {
	go func() {
		err := outboxProcessor.Start(ctx)
		if err != nil {
			logger.Error().Err(err).Msg("baskets outbox processor encountered an error")
		}
	}()
}
