package commands

import (
	"context"

	"github.com/stackus/errors"

	"github.com/LordMoMA/Intelli-Mall/depot/internal/domain"
	"github.com/LordMoMA/Intelli-Mall/internal/ddd"
)

type CreateShoppingList struct {
	ID      string
	OrderID string
	Items   []OrderItem
}

type CreateShoppingListHandler struct {
	shoppingLists   domain.ShoppingListRepository
	stores          domain.StoreRepository
	products        domain.ProductRepository
	domainPublisher ddd.EventPublisher[ddd.AggregateEvent]
}

func NewCreateShoppingListHandler(shoppingLists domain.ShoppingListRepository, stores domain.StoreRepository,
	products domain.ProductRepository, domainPublisher ddd.EventPublisher[ddd.AggregateEvent],
) CreateShoppingListHandler {
	return CreateShoppingListHandler{
		shoppingLists:   shoppingLists,
		stores:          stores,
		products:        products,
		domainPublisher: domainPublisher,
	}
}

func (h CreateShoppingListHandler) CreateShoppingList(ctx context.Context, cmd CreateShoppingList) error {
	list := domain.CreateShoppingList(cmd.ID, cmd.OrderID)

	for _, item := range cmd.Items {
		// horribly inefficient
		store, err := h.stores.Find(ctx, item.StoreID)
		if err != nil {
			return errors.Wrap(err, "building shopping list")
		}
		product, err := h.products.Find(ctx, item.ProductID)
		if err != nil {
			return errors.Wrap(err, "building shopping list")
		}
		err = list.AddItem(store, product, item.Quantity)
		if err != nil {
			return errors.Wrap(err, "building shopping list")
		}
	}

	if err := h.shoppingLists.Save(ctx, list); err != nil {
		return errors.Wrap(err, "scheduling shopping")
	}

	// publish domain events
	if err := h.domainPublisher.Publish(ctx, list.Events()...); err != nil {
		return err
	}

	return nil
}
