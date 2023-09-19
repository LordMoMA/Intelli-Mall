package queries

import (
	"context"

	"github.com/stackus/errors"

	"github.com/LordMoMA/Intelli-Mall/ordering/internal/domain"
)

type GetOrder struct {
	ID string
}

type GetOrderHandler struct {
	repo domain.OrderRepository
}

func NewGetOrderHandler(repo domain.OrderRepository) GetOrderHandler {
	return GetOrderHandler{repo: repo}
}

func (h GetOrderHandler) GetOrder(ctx context.Context, query GetOrder) (*domain.Order, error) {
	order, err := h.repo.Load(ctx, query.ID)

	return order, errors.Wrap(err, "get order query")
}
