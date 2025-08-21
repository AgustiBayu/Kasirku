package repositories

import (
	"context"
	"database/sql"
	"kasirku/models/domain"
)

type ProductRepository interface {
	Create(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Product, map[int]domain.ProductCategory)
	FindById(ctx context.Context, tx *sql.Tx, produkId int) domain.Product
	Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Delete(ctx context.Context, tx *sql.Tx, product domain.Product)
}
