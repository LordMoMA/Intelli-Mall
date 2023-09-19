package commands

import (
	"context"

	"github.com/LordMoMA/Intelli-Mall/internal/ddd"
	"github.com/LordMoMA/Intelli-Mall/stores/internal/domain"
)

type (
	CreateStore struct {
		ID       string
		Name     string
		Location string
	}

	CreateStoreHandler struct {
		stores    domain.StoreRepository
		publisher ddd.EventPublisher[ddd.Event]
	}
)

func NewCreateStoreHandler(stores domain.StoreRepository, publisher ddd.EventPublisher[ddd.Event]) CreateStoreHandler {
	return CreateStoreHandler{
		stores:    stores,
		publisher: publisher,
	}
}

func (h CreateStoreHandler) CreateStore(ctx context.Context, cmd CreateStore) error {
	store, err := h.stores.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	event, err := store.InitStore(cmd.Name, cmd.Location)
	if err != nil {
		return err
	}

	err = h.stores.Save(ctx, store)
	if err != nil {
		return err
	}

	return h.publisher.Publish(ctx, event)
}
