package application

import (
	"context"

	"github.com/LordMoMA/Intelli-Mall/notifications/internal/models"
)

type CustomerRepository interface {
	Find(ctx context.Context, customerID string) (*models.Customer, error)
}
