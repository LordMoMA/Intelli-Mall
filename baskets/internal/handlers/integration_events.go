package handlers

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/LordMoMA/Intelli-Mall/baskets/internal/domain"
	"github.com/LordMoMA/Intelli-Mall/internal/am"
	"github.com/LordMoMA/Intelli-Mall/internal/ddd"
	"github.com/LordMoMA/Intelli-Mall/internal/errorsotel"
	"github.com/LordMoMA/Intelli-Mall/internal/registry"
	"github.com/LordMoMA/Intelli-Mall/stores/storespb"
)

type integrationHandlers[T ddd.Event] struct {
	stores   domain.StoreCacheRepository
	products domain.ProductCacheRepository
}

var _ ddd.EventHandler[ddd.Event] = (*integrationHandlers[ddd.Event])(nil)

func NewIntegrationEventHandlers(reg registry.Registry, stores domain.StoreCacheRepository, products domain.ProductCacheRepository, mws ...am.MessageHandlerMiddleware) am.MessageHandler {
	return am.NewEventHandler(reg, integrationHandlers[ddd.Event]{
		stores:   stores,
		products: products,
	}, mws...)
}

func RegisterIntegrationEventHandlers(subscriber am.MessageSubscriber, handlers am.MessageHandler) (err error) {
	_, err = subscriber.Subscribe(storespb.StoreAggregateChannel, handlers, am.MessageFilter{
		storespb.StoreCreatedEvent,
		storespb.StoreRebrandedEvent,
	}, am.GroupName("baskets-stores"))
	if err != nil {
		return err
	}

	_, err = subscriber.Subscribe(storespb.ProductAggregateChannel, handlers, am.MessageFilter{
		storespb.ProductAddedEvent,
		storespb.ProductRebrandedEvent,
		storespb.ProductPriceIncreasedEvent,
		storespb.ProductPriceDecreasedEvent,
		storespb.ProductRemovedEvent,
	}, am.GroupName("baskets-products"))

	return err
}

func (h integrationHandlers[T]) HandleEvent(ctx context.Context, event T) (err error) {
	span := trace.SpanFromContext(ctx)
	defer func(started time.Time) {
		if err != nil {
			span.AddEvent(
				"Encountered an error handling integration event",
				trace.WithAttributes(errorsotel.ErrAttrs(err)...),
			)
		}
		span.AddEvent("Handled integration event", trace.WithAttributes(
			attribute.Int64("TookMS", time.Since(started).Milliseconds()),
		))
	}(time.Now())

	span.AddEvent("Handling integration event", trace.WithAttributes(
		attribute.String("Event", event.EventName()),
	))

	switch event.EventName() {
	case storespb.StoreCreatedEvent:
		return h.onStoreCreated(ctx, event)
	case storespb.StoreRebrandedEvent:
		return h.onStoreRebranded(ctx, event)
	case storespb.ProductAddedEvent:
		return h.onProductAdded(ctx, event)
	case storespb.ProductRebrandedEvent:
		return h.onProductRebranded(ctx, event)
	case storespb.ProductPriceIncreasedEvent, storespb.ProductPriceDecreasedEvent:
		return h.onProductPriceChanged(ctx, event)
	case storespb.ProductRemovedEvent:
		return h.onProductRemoved(ctx, event)
	}

	return nil
}

func (h integrationHandlers[T]) onStoreCreated(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.StoreCreated)
	return h.stores.Add(ctx, payload.GetId(), payload.GetName())
}

func (h integrationHandlers[T]) onStoreRebranded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.StoreRebranded)
	return h.stores.Rename(ctx, payload.GetId(), payload.GetName())
}

func (h integrationHandlers[T]) onProductAdded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.ProductAdded)
	return h.products.Add(ctx, payload.GetId(), payload.GetStoreId(), payload.GetName(), payload.GetPrice())
}

func (h integrationHandlers[T]) onProductRebranded(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.ProductRebranded)
	return h.products.Rebrand(ctx, payload.GetId(), payload.GetName())
}

func (h integrationHandlers[T]) onProductPriceChanged(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.ProductPriceChanged)
	return h.products.UpdatePrice(ctx, payload.GetId(), payload.GetDelta())
}

func (h integrationHandlers[T]) onProductRemoved(ctx context.Context, event ddd.Event) error {
	payload := event.Payload().(*storespb.ProductRemoved)
	return h.products.Remove(ctx, payload.GetId())
}
