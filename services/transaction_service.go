package services

import (
	"context"
	"kasirku/models/domain"
)

type TransactionService interface {
	Create(ctx context.Context, req *domain.TransactionCreateRequest) (*domain.TransactionResponse, error)
}
