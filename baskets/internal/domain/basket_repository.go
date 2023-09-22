package domain

import (
	"context"
)

/*
the Load method is not defined in the BasketRepository interface itself.
The Load method is only defined in the concrete implementations of the BasketRepository interface,
such as FakeBasketRepository.

The purpose of the BasketRepository interface is to define a set of methods
that must be implemented by any concrete type that implements the interface.
In this case, the BasketRepository interface defines the Load and Save methods,
which must be implemented by any concrete type that implements the BasketRepository interface.
*/
type BasketRepository interface {
	Load(ctx context.Context, basketID string) (*Basket, error)
	Save(ctx context.Context, basket *Basket) error
}
