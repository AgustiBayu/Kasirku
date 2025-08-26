package repositories

import (
	"context"
	"kasirku/models/domain"
)

type ProductRepository interface {
	Create(ctx context.Context, product *domain.Product) (*domain.Product, error)
	FindAll(ctx context.Context) ([]*domain.Product, map[int]*domain.ProductCategory, error)
	FindById(ctx context.Context, produkId uint) (*domain.Product, *domain.ProductCategory, error)
	Update(ctx context.Context, product *domain.Product) (*domain.Product, error)
	Delete(ctx context.Context, product *domain.Product) error
	UploadThumbnail(ctx context.Context, productId uint, path string) error
}
