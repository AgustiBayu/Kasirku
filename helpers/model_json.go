package helpers

import (
	"kasirku/models/domain"
)

func ToProductCategoryResponse(category *domain.ProductCategory) *domain.ProductCategoryResponse {
	return &domain.ProductCategoryResponse{
		ID:       category.ID,
		Category: category.Category,
	}
}
func ToProductCategoryResponses(categories []*domain.ProductCategory) []*domain.ProductCategoryResponse {
	var responses []*domain.ProductCategoryResponse
	for _, category := range categories {
		responses = append(responses, ToProductCategoryResponse(category))
	}
	return responses
}
func ToProductResponse(product *domain.Product, category *domain.ProductCategory) *domain.ProductResponse {
	return &domain.ProductResponse{
		ID:         product.ID,
		Name:       product.Name,
		Slug:       product.Slug,
		Thumbnail:  product.Thumbnail,
		Price:      product.Price,
		Exp:        FormatTanggal(product.Exp),
		Stock:      product.Stock,
		CategoryID: product.CategoryID,
		ProductCategory: domain.ProductCategoryResponse{
			ID:       category.ID,
			Category: category.Category,
		},
		Barcode: product.Barcode,
	}
}
func ToProductResponses(producs []*domain.Product, categoriesMap map[uint]*domain.ProductCategory) []*domain.ProductResponse {
	var productResponses []*domain.ProductResponse
	for _, product := range producs {
		category, exits := categoriesMap[product.CategoryID]
		if !exits {
			category = &domain.ProductCategory{}
		}
		productResponses = append(productResponses, ToProductResponse(product, category))
	}
	return productResponses
}

func ToTransactionResponse(transaction *domain.Transaction, details []*domain.TransactionDetail) *domain.TransactionResponse {
	var detailResponses []domain.TransactionDetailResponse
	for _, d := range details {
		detailResponses = append(detailResponses, domain.TransactionDetailResponse{
			ProductID:          int(d.ProductID),
			Quantity:           int(d.Quantity),
			PriceAtTransaction: int(d.PriceAtTransaction),
			Subtotal:           int(d.Subtotal),
		})
	}

	return &domain.TransactionResponse{
		ID:            int(transaction.ID),
		CreatedAt:     FormatTanggal(transaction.CreatedAt),
		TotalAmount:   int(transaction.TotalAmount),
		PaymentMethod: transaction.PaymentMethod,
		AmountPaid:    int(transaction.AmountPaid),
		Change:        int(transaction.Change),
		Details:       detailResponses,
	}
}
