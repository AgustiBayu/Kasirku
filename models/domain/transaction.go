package domain

import "time"

type Transaction struct {
	ID            uint      `gorm:"primaryKey"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	TotalAmount   uint
	PaymentMethod string
	AmountPaid    uint
	Change        uint
}

type TransactionCreateRequest struct {
	Items         []TransactionItemRequest `json:"items" validate:"required,min=1"`
	PaymentMethod string                   `json:"payment_method" validate:"required"`
	AmountPaid    uint                     `json:"amount_paid" validate:"required,gte=0"`
}

type TransactionResponse struct {
	ID            int                         `json:"id"`
	CreatedAt     string                      `json:"created_at"`
	TotalAmount   int                         `json:"total_amount"`
	PaymentMethod string                      `json:"payment_method"`
	AmountPaid    int                         `json:"amount_paid"`
	Change        int                         `json:"change"`
	Details       []TransactionDetailResponse `json:"details"`
}
