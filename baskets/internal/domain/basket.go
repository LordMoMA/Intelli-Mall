package domain

import (
	"github.com/stackus/errors"

	"github.com/LordMoMA/Intelli-Mall/internal/ddd"
	"github.com/LordMoMA/Intelli-Mall/internal/es"
)

const BasketAggregate = "baskets.Basket"

var (
	ErrBasketHasNoItems         = errors.Wrap(errors.ErrBadRequest, "the basket has no items")
	ErrBasketCannotBeModified   = errors.Wrap(errors.ErrBadRequest, "the basket cannot be modified")
	ErrBasketCannotBeCancelled  = errors.Wrap(errors.ErrBadRequest, "the basket cannot be cancelled")
	ErrQuantityCannotBeNegative = errors.Wrap(errors.ErrBadRequest, "the item quantity cannot be negative")
	ErrBasketIDCannotBeBlank    = errors.Wrap(errors.ErrBadRequest, "the basket id cannot be blank")
	ErrPaymentIDCannotBeBlank   = errors.Wrap(errors.ErrBadRequest, "the payment id cannot be blank")
	ErrCustomerIDCannotBeBlank  = errors.Wrap(errors.ErrBadRequest, "the customer id cannot be blank")
)

type Basket struct {
	es.Aggregate
	CustomerID string
	PaymentID  string
	Items      map[string]Item
	Status     BasketStatus
}

var _ interface {
	es.EventApplier
	es.Snapshotter
} = (*Basket)(nil)

func NewBasket(id string) *Basket {
	return &Basket{
		Aggregate: es.NewAggregate(id, BasketAggregate),
		Items:     make(map[string]Item),
	}
}

func (b *Basket) Start(customerID string) (ddd.Event, error) {
	if b.Status != BasketUnknown {
		return nil, ErrBasketCannotBeModified
	}

	if customerID == "" {
		return nil, ErrCustomerIDCannotBeBlank
	}

	b.AddEvent(BasketStartedEvent, &BasketStarted{
		CustomerID: customerID,
	})

	return ddd.NewEvent(BasketStartedEvent, b), nil
}

func (b Basket) IsCancellable() bool {
	return b.Status == BasketIsOpen
}

func (b Basket) IsOpen() bool {
	return b.Status == BasketIsOpen
}

func (b *Basket) Cancel() (ddd.Event, error) {
	if !b.IsCancellable() {
		return nil, ErrBasketCannotBeCancelled
	}

	b.AddEvent(BasketCanceledEvent, &BasketCanceled{})

	return ddd.NewEvent(BasketCanceledEvent, b), nil
}

func (b *Basket) Checkout(paymentID string) (ddd.Event, error) {
	if !b.IsOpen() {
		return nil, ErrBasketCannotBeModified
	}

	if len(b.Items) == 0 {
		return nil, ErrBasketHasNoItems
	}

	if paymentID == "" {
		return nil, ErrPaymentIDCannotBeBlank
	}

	b.AddEvent(BasketCheckedOutEvent, &BasketCheckedOut{
		PaymentID: paymentID,
	})

	return ddd.NewEvent(BasketCheckedOutEvent, b), nil
}

func (b *Basket) AddItem(store *Store, product *Product, quantity int) error {
	if !b.IsOpen() {
		return ErrBasketCannotBeModified
	}

	if quantity < 0 {
		return ErrQuantityCannotBeNegative
	}

	b.AddEvent(BasketItemAddedEvent, &BasketItemAdded{
		Item: Item{
			StoreID:      store.ID,
			ProductID:    product.ID,
			StoreName:    store.Name,
			ProductName:  product.Name,
			ProductPrice: product.Price,
			Quantity:     quantity,
		},
	})

	return nil
}

func (b *Basket) RemoveItem(product *Product, quantity int) error {
	if !b.IsOpen() {
		return ErrBasketCannotBeModified
	}

	if quantity < 0 {
		return ErrQuantityCannotBeNegative
	}

	if _, exists := b.Items[product.ID]; exists {
		b.AddEvent(BasketItemRemovedEvent, &BasketItemRemoved{
			ProductID: product.ID,
			Quantity:  quantity,
		})
	}

	return nil
}

func (b *Basket) ApplyEvent(event ddd.Event) error {
	switch payload := event.Payload().(type) {
	case *BasketStarted:
		b.CustomerID = payload.CustomerID
		b.Status = BasketIsOpen

	case *BasketItemAdded:
		if item, exists := b.Items[payload.Item.ProductID]; exists {
			item.Quantity += payload.Item.Quantity
			b.Items[payload.Item.ProductID] = item
		} else {
			b.Items[payload.Item.ProductID] = payload.Item
		}

	case *BasketItemRemoved:
		if item, exists := b.Items[payload.ProductID]; exists {
			if item.Quantity-payload.Quantity <= 1 {
				delete(b.Items, payload.ProductID)
			} else {
				item.Quantity -= payload.Quantity
				b.Items[payload.ProductID] = item
			}
		}

	case *BasketCanceled:
		b.Items = make(map[string]Item)
		b.Status = BasketIsCanceled

	case *BasketCheckedOut:
		b.PaymentID = payload.PaymentID
		b.Status = BasketIsCheckedOut

	default:
		return errors.ErrInternal.Msgf("%T received the event %s with unexpected payload %T", b, event.EventName(), payload)
	}

	return nil
}

func (b *Basket) ApplySnapshot(snapshot es.Snapshot) error {
	switch ss := snapshot.(type) {
	case *BasketV1:
		b.CustomerID = ss.CustomerID
		b.PaymentID = ss.PaymentID
		b.Items = ss.Items
		b.Status = ss.Status

	default:
		return errors.ErrInternal.Msgf("%T received the unexpected snapshot %T", b, snapshot)
	}

	return nil
}

func (b *Basket) ToSnapshot() es.Snapshot {
	return &BasketV1{
		CustomerID: b.CustomerID,
		PaymentID:  b.PaymentID,
		Items:      b.Items,
		Status:     b.Status,
	}
}
