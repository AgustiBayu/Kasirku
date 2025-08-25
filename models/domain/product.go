package domain

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Slug       string
	Thumbnail  string
	Price      int
	Exp        time.Time
	CategoryId int
	DeletedAt  gorm.DeletedAt  `gorm:"index"`
	Category   ProductCategory `gorm:"foreignKey:CategoryId"`
}

type ProductCreateRequest struct {
	ID              uint                    `validate:"required" json:"id"`
	Name            string                  `validate:"required" json:"name"`
	Slug            string                  `validate:"required" json:"slug"`
	Thumbnail       string                  `validate:"required" json:"thumbnail"`
	Price           int                     `validate:"required" json:"price"`
	Exp             string                  `validate:"required" json:"exp"`
	ProductCategory ProductCategoryResponse `validate:"required" json:"product_category"`
}

type ProductResponse struct {
	ID              uint                    `json:"id"`
	Name            string                  `json:"name"`
	Slug            string                  `json:"slug"`
	Thumbnail       string                  `json:"thumbnail"`
	Price           int                     `json:"price"`
	Exp             string                  `json:"exp"`
	ProductCategory ProductCategoryResponse `json:"product_category"`
}

type ProductUpdateRequest struct {
	ID              uint                    `validate:"required" json:"id"`
	Name            string                  `validate:"required" json:"name"`
	Slug            string                  `validate:"required" json:"slug"`
	Thumbnail       string                  `validate:"required" json:"thumbnail"`
	Price           int                     `validate:"required" json:"price"`
	Exp             string                  `validate:"required" json:"exp"`
	ProductCategory ProductCategoryResponse `validate:"required" json:"product_category"`
}
