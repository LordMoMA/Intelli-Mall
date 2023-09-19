package storespb

import (
	"github.com/LordMoMA/Intelli-Mall/internal/registry"
	"github.com/LordMoMA/Intelli-Mall/internal/registry/serdes"
)

const (
	StoreAggregateChannel = "intellimall.stores.events.Store"

	StoreCreatedEvent              = "storesapi.StoreCreated"
	StoreParticipatingToggledEvent = "storesapi.StoreParticipatingToggled"
	StoreRebrandedEvent            = "storesapi.StoreRebranded"

	ProductAggregateChannel = "intellimall.stores.events.Product"

	ProductAddedEvent          = "storesapi.ProductAdded"
	ProductRebrandedEvent      = "storesapi.ProductRebranded"
	ProductPriceIncreasedEvent = "storesapi.ProductPriceIncreased"
	ProductPriceDecreasedEvent = "storesapi.ProductPriceDecreased"
	ProductRemovedEvent        = "storesapi.ProductRemoved"
)

func Registrations(reg registry.Registry) error {
	return RegistrationsWithSerde(serdes.NewProtoSerde(reg))
}

func RegistrationsWithSerde(serde registry.Serde) error {
	// Store events
	if err := serde.Register(&StoreCreated{}); err != nil {
		return err
	}
	if err := serde.Register(&StoreParticipationToggled{}); err != nil {
		return err
	}
	if err := serde.Register(&StoreRebranded{}); err != nil {
		return err
	}

	if err := serde.Register(&ProductAdded{}); err != nil {
		return err
	}
	if err := serde.Register(&ProductRebranded{}); err != nil {
		return err
	}
	if err := serde.RegisterKey(ProductPriceIncreasedEvent, &ProductPriceChanged{}); err != nil {
		return err
	}
	if err := serde.RegisterKey(ProductPriceDecreasedEvent, &ProductPriceChanged{}); err != nil {
		return err
	}
	if err := serde.Register(&ProductRemoved{}); err != nil {
		return err
	}

	return nil
}

func (*StoreCreated) Key() string              { return StoreCreatedEvent }
func (*StoreParticipationToggled) Key() string { return StoreParticipatingToggledEvent }
func (*StoreRebranded) Key() string            { return StoreRebrandedEvent }

func (*ProductAdded) Key() string     { return ProductAddedEvent }
func (*ProductRebranded) Key() string { return ProductRebrandedEvent }
func (*ProductRemoved) Key() string   { return ProductRemovedEvent }
