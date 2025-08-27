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
		ID:        product.ID,
		Name:      product.Name,
		Slug:      product.Slug,
		Thumbnail: product.Thumbnail,
		Price:     product.Price,
		Exp:       FormatTanggal(product.Exp),
		ProductCategory: domain.ProductCategoryResponse{
			ID:       category.ID,
			Category: category.Category,
		},
	}
}
func ToProductResponses(producs []*domain.Product, categoriesMap map[int]*domain.ProductCategory) []*domain.ProductResponse {
	var productResponses []*domain.ProductResponse
	for _, product := range producs {
		category, exits := categoriesMap[product.CategoryId]
		if !exits {
			category = &domain.ProductCategory{}
		}
		productResponses = append(productResponses, ToProductResponse(product, category))
	}
	return productResponses
}
