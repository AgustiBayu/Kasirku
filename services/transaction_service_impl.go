package services

import (
	"context"
	"kasirku/exception"
	"kasirku/helpers"
	"kasirku/models/domain"
	"kasirku/repositories"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type TransactionServiceImpl struct {
	TransactionRepo repositories.TransactionRepository
	ProductRepo     repositories.ProductRepository
	DB              *gorm.DB
	Validate        *validator.Validate
}

func NewTransactionService(transactionRepo repositories.TransactionRepository, productRepo repositories.ProductRepository, db *gorm.DB, validate *validator.Validate) TransactionService {
	return &TransactionServiceImpl{
		TransactionRepo: transactionRepo,
		ProductRepo:     productRepo,
		DB:              db,
		Validate:        validate,
	}
}

func (s *TransactionServiceImpl) Create(ctx context.Context, req *domain.TransactionCreateRequest) (*domain.TransactionResponse, error) {
	if err := s.Validate.Struct(req); err != nil {
		return nil, exception.BadRequest("field not found")
	}

	var details []*domain.TransactionDetail
	var savedTransaction *domain.Transaction

	err := s.DB.Transaction(func(tx *gorm.DB) error {
		var totalAmount uint = 0
		// 1. Validate items and check stock
		for _, item := range req.Items {
			product, _, err := s.ProductRepo.FindById(ctx, item.ProductID)
			if err != nil {
				return exception.BadRequest("product not found")
			}

			// Check if product stock is sufficient
			if product.Stock < item.Quantity {
				return exception.BadRequest("Insufficient stock for product: " + product.Name)
			}

			subtotal := product.Price * item.Quantity
			totalAmount += subtotal

			details = append(details, &domain.TransactionDetail{
				ProductID:          product.ID,
				Quantity:           item.Quantity,
				PriceAtTransaction: product.Price,
				Subtotal:           subtotal,
			})
		}

		if req.AmountPaid < totalAmount {
			return exception.InternalServerError("insufficient payment amount")
		}

		change := req.AmountPaid - totalAmount

		transaction := &domain.Transaction{
			TotalAmount:   totalAmount,
			PaymentMethod: req.PaymentMethod,
			AmountPaid:    req.AmountPaid,
			Change:        change,
		}

		// 2. Create transaction record
		var createErr error
		savedTransaction, createErr = s.TransactionRepo.Create(ctx, transaction, details)
		if createErr != nil {
			return exception.InternalServerError("transaction failed")
		}

		// 3. Decrease product stock after successful transaction
		for _, item := range details {
			product, _, err := s.ProductRepo.FindById(ctx, item.ProductID)
			if err != nil {
				return exception.InternalServerError("product not found during stock update")
			}
			newStock := product.Stock - item.Quantity
			if err := s.ProductRepo.UpdateStock(ctx, product.ID, newStock); err != nil {
				return exception.InternalServerError("failed to update stock")
			}
		}

		return nil // Commit transaction
	})

	if err != nil {
		return nil, err
	}

	return helpers.ToTransactionResponse(savedTransaction, details), nil
}
