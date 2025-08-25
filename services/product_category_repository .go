package repositories

import (
	"context"
	"kasirku/models/domain"
)

type ProductCategoryService interface {
	Create(ctx context.Context, req *domain.ProductCategoryCreateRequest) error
	FindAll(ctx context.Context) ([]*domain.ProductCategoryResponse, error)
	FindById(ctx context.Context, produkId int) (*domain.ProductCategoryResponse, error)
	Update(ctx context.Context, req *domain.ProductCategoryUpdateRequest) error
	Delete(ctx context.Context, produkId int) error
}
