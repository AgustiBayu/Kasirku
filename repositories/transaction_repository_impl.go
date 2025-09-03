package repositories

import (
	"context"
	"kasirku/models/domain"

	"gorm.io/gorm"
)

type TransactionRepositoryImpl struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &TransactionRepositoryImpl{
		DB: db,
	}
}

func (r *TransactionRepositoryImpl) Create(ctx context.Context, transaction *domain.Transaction, details []*domain.TransactionDetail) (*domain.Transaction, error) {
	err := r.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 1. Create the main transaction record
		if err := tx.Create(transaction).Error; err != nil {
			return err
		}

		// 2. Associate details with the new transaction ID and create them
		for i := range details {
			details[i].TransactionID = transaction.ID
		}

		if err := tx.Create(&details).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return transaction, nil
}
