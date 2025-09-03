package repositories

import (
	"context"
	"kasirku/models/domain"
)

type TransactionRepository interface {
	Create(ctx context.Context, transaction *domain.Transaction, details []*domain.TransactionDetail) (*domain.Transaction, error)
}
