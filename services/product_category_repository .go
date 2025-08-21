package repositories

import (
	"context"
	"database/sql"
	"kasirku/models/domain"
)

type ProductCategoryRepository interface {
	Create(ctx context.Context, tx *sql.Tx, categoryProduct domain.ProductCategory) domain.ProductCategory
	FindAll(ctx context.Context, tx *sql.Tx) []domain.ProductCategory
	FindById(ctx context.Context, tx *sql.Tx, categoryProductId int) domain.ProductCategory
	Update(ctx context.Context, tx *sql.Tx, categoryProduct domain.ProductCategory) domain.ProductCategory
	Delete(ctx context.Context, tx *sql.Tx, categoryProduct domain.ProductCategory)
}
