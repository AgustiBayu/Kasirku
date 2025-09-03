package domain

type TransactionDetail struct {
	ID                 uint `gorm:"primaryKey"`
	TransactionID      uint
	Transaction        Transaction
	ProductID          uint
	Product            Product
	Quantity           uint
	PriceAtTransaction uint
	Subtotal           uint
}

type TransactionItemRequest struct {
	ProductID uint `json:"product_id" validate:"required"`
	Quantity  uint `json:"quantity" validate:"required,gt=0"`
}

type TransactionDetailResponse struct {
	ProductID          int `json:"product_id"`
	Quantity           int `json:"quantity"`
	PriceAtTransaction int `json:"price_at_transaction"`
	Subtotal           int `json:"subtotal"`
}
