package repositories

import (
	"context"
	"kasirku/models/web"
)

type ProductService interface {
	Save(ctx context.Context, request web.ProductCategoryCreateRequest) error
	FindAll(ctx context.Context) []web.ProductCategoryResponse
	FindById(ctx context.Context, produkId int) web.ProductCategoryResponse
	Update(ctx context.Context, request web.ProductCategoryUpdateRequest) error
	Destroy(ctx context.Context, produkId int) error
}
