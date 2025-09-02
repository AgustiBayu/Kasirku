package services

import (
	"context"
	"io"
	"kasirku/exception"
	"kasirku/helpers"
	"kasirku/models/domain"
	"kasirku/repositories"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductRepository         repositories.ProductRepository
	ProductCategoryRepository repositories.ProductCategoryRepository
	Validate                  *validator.Validate
}

func NewProductService(productRepository repositories.ProductRepository, productCategoryRepository repositories.ProductCategoryRepository,
	validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository:         productRepository,
		ProductCategoryRepository: productCategoryRepository,
		Validate:                  validate,
	}
}

func (p *ProductServiceImpl) Create(ctx context.Context, req *domain.ProductCreateRequest) error {
	if err := p.Validate.Struct(req); err != nil {
		return exception.BadRequest("field is not falid")
	}
	EXP, err := time.Parse("2006-01-02", req.Exp)
	if err != nil {
		return exception.BadRequest("invalid date format, use dd-mm-yyyy")
	}
	product := domain.Product{
		Name:       req.Name,
		Slug:       req.Slug,
		Thumbnail:  "",
		Price:      req.Price,
		Exp:        EXP,
		CategoryID: req.CategoryID,
	}
	if _, err := p.ProductRepository.Create(ctx, &product); err != nil {
		return exception.InternalServerError("failed to create product")
	}
	return nil
}
func (p *ProductServiceImpl) FindAll(ctx context.Context) ([]*domain.ProductResponse, error) {
	product, category, err := p.ProductRepository.FindAll(ctx)
	if err != nil {
		return nil, exception.InternalServerError("data is not found")
	}
	if len(product) == 0 {
		return []*domain.ProductResponse{}, nil
	}
	return helpers.ToProductResponses(product, category), nil
}
func (p *ProductServiceImpl) FindById(ctx context.Context, produkId int) (*domain.ProductResponse, error) {
	product, category, err := p.ProductRepository.FindById(ctx, uint(produkId))
	if err != nil {
		return nil, exception.NotFound("id category not exists")
	}
	return helpers.ToProductResponse(product, category), nil
}
func (p *ProductServiceImpl) Update(ctx context.Context, req *domain.ProductUpdateRequest) error {
	if err := p.Validate.Struct(req); err != nil {
		return exception.BadRequest("field is not falid")
	}
	product, _, err := p.ProductRepository.FindById(ctx, uint(req.ID))
	if err != nil {
		return exception.NotFound("id category not exists")
	}
	EXP, err := time.Parse("2006-01-02", req.Exp)
	if err != nil {
		return exception.BadRequest("invalid date format, use dd-mm-yyyy")
	}
	product.Name = req.Name
	product.Slug = req.Slug
	product.Price = req.Price
	product.Exp = EXP
	product.CategoryID = req.CategoryID
	if _, err := p.ProductRepository.Update(ctx, product); err != nil {
		return exception.InternalServerError("failed to update product")
	}
	return nil
}
func (p *ProductServiceImpl) Delete(ctx context.Context, produkId int) error {
	product, _, err := p.ProductRepository.FindById(ctx, uint(produkId))
	if err != nil {
		return exception.NotFound("id product not exist")
	}
	if err := p.ProductRepository.Delete(ctx, product); err != nil {
		return exception.InternalServerError("failed to delete product")
	}
	return nil
}
func (p *ProductServiceImpl) UploadThumbnail(ctx context.Context, productId uint, file multipart.File, handler *multipart.FileHeader) error {
	defer file.Close()
	product, _, err := p.ProductRepository.FindById(ctx, productId)
	if err != nil {
		return exception.NotFound("id product not exists")
	}
	if err := os.MkdirAll("images", os.ModePerm); err != nil {
		return exception.InternalServerError("failed to create image directory")
	}
	filePath := filepath.Join("images", handler.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		return exception.InternalServerError("failed to save image")
	}
	defer dst.Close()
	if _, err := io.Copy(dst, file); err != nil {
		return exception.InternalServerError("failed to copy image file")
	}
	if err := p.ProductRepository.UploadThumbnail(ctx, product.ID, filePath); err != nil {
		return exception.InternalServerError("failed to update product thumbnail")
	}
	return nil
}
