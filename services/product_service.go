package services

import (
	"context"
	"kasirku/models/domain"
	"mime/multipart"
)

type ProductService interface {
	Create(ctx context.Context, req *domain.ProductCreateRequest) error
	FindAll(ctx context.Context) ([]*domain.ProductResponse, error)
	FindById(ctx context.Context, produkId int) (*domain.ProductResponse, error)
	FindByBarcode(ctx context.Context, barcode string) (*domain.ProductResponse, error)
	Update(ctx context.Context, req *domain.ProductUpdateRequest) error
	Delete(ctx context.Context, produkId int) error
	UploadThumbnail(ctx context.Context, productId uint, file multipart.File, handler *multipart.FileHeader) error
}
