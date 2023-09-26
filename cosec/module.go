package cosec

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog"

	"github.com/LordMoMA/Intelli-Mall/cosec/internal"
	"github.com/LordMoMA/Intelli-Mall/cosec/internal/constants"
	"github.com/LordMoMA/Intelli-Mall/cosec/internal/handlers"
	"github.com/LordMoMA/Intelli-Mall/cosec/internal/models"
	"github.com/LordMoMA/Intelli-Mall/customers/customerspb"
	"github.com/LordMoMA/Intelli-Mall/depot/depotpb"
	"github.com/LordMoMA/Intelli-Mall/internal/am"
	"github.com/LordMoMA/Intelli-Mall/internal/amotel"
	"github.com/LordMoMA/Intelli-Mall/internal/amprom"
	"github.com/LordMoMA/Intelli-Mall/internal/di"
	"github.com/LordMoMA/Intelli-Mall/internal/jetstream"
	pg "github.com/LordMoMA/Intelli-Mall/internal/postgres"
	"github.com/LordMoMA/Intelli-Mall/internal/postgresotel"
	"github.com/LordMoMA/Intelli-Mall/internal/registry"
	"github.com/LordMoMA/Intelli-Mall/internal/registry/serdes"
	"github.com/LordMoMA/Intelli-Mall/internal/sec"
	"github.com/LordMoMA/Intelli-Mall/internal/system"
	"github.com/LordMoMA/Intelli-Mall/internal/tm"
	"github.com/LordMoMA/Intelli-Mall/ordering/orderingpb"
	"github.com/LordMoMA/Intelli-Mall/payments/paymentspb"
)

// cosec is short for Create-Order-Saga-Execution-Coordinator

type Module struct{}

func (Module) Startup(ctx context.Context, mono system.Service) (err error) {
	return Root(ctx, mono)
}

func Root(ctx context.Context, svc system.Service) (err error) {
	container := di.New()
	// setup Driven adapters
	container.AddSingleton(constants.RegistryKey, func(c di.Container) (any, error) {
		reg := registry.New()
		if err := registrations(reg); err != nil {
			return nil, err
		}
		if err := orderingpb.Registrations(reg); err != nil {
			return nil, err
		}
		if err := customerspb.Registrations(reg); err != nil {
			return nil, err
		}
		if err := depotpb.Registrations(reg); err != nil {
			return nil, err
		}
		if err := paymentspb.Registrations(reg); err != nil {
			return nil, err
		}
		return reg, nil
	})
	stream := jetstream.NewStream(svc.Config().Nats.Stream, svc.JS(), svc.Logger())
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
	container.AddScoped(constants.CommandPublisherKey, func(c di.Container) (any, error) {
		return am.NewCommandPublisher(
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.MessagePublisherKey).(am.MessagePublisher),
		), nil
	})
	container.AddScoped(constants.InboxStoreKey, func(c di.Container) (any, error) {
		tx := postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx))
		return pg.NewInboxStore(constants.InboxTableName, tx), nil
	})
	container.AddScoped(constants.SagaStoreKey, func(c di.Container) (any, error) {
		reg := c.Get(constants.RegistryKey).(registry.Registry)
		return sec.NewSagaRepository[*models.CreateOrderData](
			reg,
			pg.NewSagaStore(
				constants.SagasTableName,
				postgresotel.Trace(c.Get(constants.DatabaseTransactionKey).(*sql.Tx)),
				reg,
			),
		), nil
	})
	container.AddSingleton(constants.SagaKey, func(c di.Container) (any, error) {
		return internal.NewCreateOrderSaga(), nil
	})

	// setup application
	container.AddScoped(constants.OrchestratorKey, func(c di.Container) (any, error) {
		return sec.NewOrchestrator[*models.CreateOrderData](
			c.Get(constants.SagaKey).(sec.Saga[*models.CreateOrderData]),
			c.Get(constants.SagaStoreKey).(sec.SagaRepository[*models.CreateOrderData]),
			c.Get(constants.CommandPublisherKey).(am.CommandPublisher),
		), nil
	})
	container.AddScoped(constants.IntegrationEventHandlersKey, func(c di.Container) (any, error) {
		return handlers.NewIntegrationEventHandlers(
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.OrchestratorKey).(sec.Orchestrator[*models.CreateOrderData]),
			tm.InboxHandler(c.Get(constants.InboxStoreKey).(tm.InboxStore)),
		), nil
	})
	container.AddScoped(constants.ReplyHandlersKey, func(c di.Container) (any, error) {
		return handlers.NewReplyHandlers(
			c.Get(constants.RegistryKey).(registry.Registry),
			c.Get(constants.OrchestratorKey).(sec.Orchestrator[*models.CreateOrderData]),
			tm.InboxHandler(c.Get(constants.InboxStoreKey).(tm.InboxStore)),
		), nil
	})
	outboxProcessor := tm.NewOutboxProcessor(
		stream,
		pg.NewOutboxStore(constants.OutboxTableName, svc.DB()),
	)

	// setup Driver adapters
	if err = handlers.RegisterIntegrationEventHandlersTx(container); err != nil {
		return err
	}
	if err = handlers.RegisterReplyHandlersTx(container); err != nil {
		return err
	}
	startOutboxProcessor(ctx, outboxProcessor, svc.Logger())

	return
}

func registrations(reg registry.Registry) (err error) {
	serde := serdes.NewJsonSerde(reg)

	// Saga data
	if err = serde.RegisterKey(internal.CreateOrderSagaName, models.CreateOrderData{}); err != nil {
		return err
	}

	return nil
}

func startOutboxProcessor(ctx context.Context, outboxProcessor tm.OutboxProcessor, logger zerolog.Logger) {
	go func() {
		err := outboxProcessor.Start(ctx)
		if err != nil {
			logger.Error().Err(err).Msg("cosec outbox processor encountered an error")
		}
	}()
}
