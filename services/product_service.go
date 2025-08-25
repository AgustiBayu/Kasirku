package repositories

import (
	"context"
	"kasirku/models/domain"
)

type ProductService interface {
	Create(ctx context.Context, req domain.ProductCreateRequest) error
	FindAll(ctx context.Context) ([]domain.ProductResponse, error)
	FindById(ctx context.Context, produkId int) (domain.ProductResponse, error)
	Update(ctx context.Context, req domain.ProductCategoryUpdateRequest) error
	Delete(ctx context.Context, produkId int) error
}
