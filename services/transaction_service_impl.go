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
	var totalAmount uint = 0

	for _, item := range req.Items {
		product, _, err := s.ProductRepo.FindById(ctx, uint(item.ProductID))
		if err != nil {
			return nil, exception.BadRequest("product not found")
		}

		// TODO: Check if product stock is sufficient

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
		return nil, exception.InternalServerError("insufficient payment amount")
	}

	change := req.AmountPaid - totalAmount

	transaction := &domain.Transaction{
		TotalAmount:   uint(totalAmount),
		PaymentMethod: req.PaymentMethod,
		AmountPaid:    req.AmountPaid,
		Change:        change,
	}

	// The repository handles the DB transaction
	savedTransaction, err := s.TransactionRepo.Create(ctx, transaction, details)
	if err != nil {
		return nil, exception.InternalServerError("transcation failed")
	}

	// TODO: Decrease product stock after successful transaction

	return helpers.ToTransactionResponse(savedTransaction, details), nil
}
