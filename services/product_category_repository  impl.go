package repositories

import (
	"context"
	"kasirku/exception"
	"kasirku/models/domain"
	"kasirku/repositories"

	"github.com/go-playground/validator/v10"
)

type ProductCategoryServiceImpl struct {
	ProductCategoryRepository repositories.ProductCategoryRepository
	Validate                  *validator.Validate
}

func NewProductCategoryRepository(productCategoryRepository repositories.ProductCategoryRepository,
	validate *validator.Validate) ProductCategoryService {
	return &ProductCategoryServiceImpl{
		ProductCategoryRepository: productCategoryRepository,
		Validate:                  validate,
	}
}

func (p *ProductCategoryServiceImpl) Create(ctx context.Context, req *domain.ProductCategoryCreateRequest) error {
	if err := p.Validate.Struct(req); err != nil {
		return exception.BadRequest("field is not valid")
	}

	category := domain.ProductCategory{
		Category: req.Category,
	}
	if _, err := p.ProductCategoryRepository.Create(ctx, &category); err != nil {
		return exception.InternalServerError("failed create data product category")
	}
	return nil
}
func (p *ProductCategoryServiceImpl) FindAll(ctx context.Context) ([]*domain.ProductCategoryResponse, error) {
	
}
func (p *ProductCategoryServiceImpl) FindById(ctx context.Context, produkId int) (*domain.ProductCategoryResponse, error) {

}
func (p *ProductCategoryServiceImpl) Update(ctx context.Context, req *domain.ProductCategoryUpdateRequest) error {

}
func (p *ProductCategoryServiceImpl) Delete(ctx context.Context, produkId int) error {

}
