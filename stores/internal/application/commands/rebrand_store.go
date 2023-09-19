package commands

import (
	"context"

	"github.com/LordMoMA/Intelli-Mall/internal/ddd"
	"github.com/LordMoMA/Intelli-Mall/stores/internal/domain"
)

type RebrandStore struct {
	ID   string
	Name string
}

type RebrandStoreHandler struct {
	stores    domain.StoreRepository
	publisher ddd.EventPublisher[ddd.Event]
}

func NewRebrandStoreHandler(stores domain.StoreRepository, publisher ddd.EventPublisher[ddd.Event]) RebrandStoreHandler {
	return RebrandStoreHandler{
		stores:    stores,
		publisher: publisher,
	}
}

func (h RebrandStoreHandler) RebrandStore(ctx context.Context, cmd RebrandStore) error {
	store, err := h.stores.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	event, err := store.Rebrand(cmd.Name)
	if err != nil {
		return err
	}

	err = h.stores.Save(ctx, store)
	if err != nil {
		return err
	}

	return h.publisher.Publish(ctx, event)
}
