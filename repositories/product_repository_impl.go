package repositories

import (
	"context"
	"kasirku/models/domain"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		DB: db,
	}
}

func (p *ProductRepositoryImpl) Create(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	if err := p.DB.WithContext(ctx).Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
func (p *ProductRepositoryImpl) FindAll(ctx context.Context) ([]*domain.Product, map[uint]*domain.ProductCategory, error) {
	var products []*domain.Product
	if err := p.DB.WithContext(ctx).Preload("Category").Find(&products).Error; err != nil {
		return nil, nil, err
	}
	categories := make(map[uint]*domain.ProductCategory)
	for i := range products {
		if products[i].Category.ID != 0 {
			categories[uint(products[i].Category.ID)] = &products[i].Category
		}
	}
	return products, categories, nil
}
func (p *ProductRepositoryImpl) FindById(ctx context.Context, produkId uint) (*domain.Product, *domain.ProductCategory, error) {
	var product domain.Product
	if err := p.DB.WithContext(ctx).Preload("Category").First(&product, produkId).Error; err != nil {
		return nil, nil, err

	}
	return &product, &product.Category, nil
}
func (p *ProductRepositoryImpl) Update(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	if err := p.DB.WithContext(ctx).Save(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
func (p *ProductRepositoryImpl) Delete(ctx context.Context, product *domain.Product) error {
	if err := p.DB.WithContext(ctx).Delete(product).Error; err != nil {
		return err
	}
	return nil
}
func (p *ProductRepositoryImpl) UploadThumbnail(ctx context.Context, productId uint, path string) error {
	var product *domain.Product
	if err := p.DB.WithContext(ctx).First(&product, productId).Error; err != nil {
		return err
	}
	product.Thumbnail = path
	if err := p.DB.WithContext(ctx).Save(&product).Error; err != nil {
		return err
	}
	return nil
}
