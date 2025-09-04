package domain

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID         uint `gorm:"primaryKey"`
	Name       string
	Slug       string
	Barcode    string `gorm:"unique"`
	Thumbnail  string
	Price      uint
	Exp        time.Time
	Stock      uint
	CategoryID uint
	DeletedAt  gorm.DeletedAt  `gorm:"index"`
	Category   ProductCategory `gorm:"foreignKey:CategoryID"`
}

type ProductCreateRequest struct {
	Name       string `validate:"required" json:"name"`
	Slug       string `validate:"required" json:"slug"`
	Barcode    string `validate:"required" json:"barcode"`
	Thumbnail  string `json:"thumbnail"`
	Price      uint   `validate:"required" json:"price"`
	Exp        string `validate:"required" json:"exp"`
	Stock      uint   `validate:"required" json:"stock"`
	CategoryID uint   `validate:"required" json:"category_id"`
}

type ProductResponse struct {
	ID              uint                    `json:"id"`
	Name            string                  `json:"name"`
	Slug            string                  `json:"slug"`
	Barcode         string                  `json:"barcode"`
	Thumbnail       string                  `json:"thumbnail"`
	Price           uint                    `json:"price"`
	Exp             string                  `json:"exp"`
	Stock           uint                    `json:"stock"`
	CategoryID      uint                    `json:"category_id"`
	ProductCategory ProductCategoryResponse `json:"category"`
}

type ProductUpdateRequest struct {
	ID         uint   `validate:"required" json:"id"`
	Name       string `validate:"required" json:"name"`
	Slug       string `validate:"required" json:"slug"`
	Barcode    string `validate:"required" json:"barcode"`
	Thumbnail  string `json:"thumbnail"`
	Price      uint   `validate:"required" json:"price"`
	Exp        string `validate:"required" json:"exp"`
	Stock      uint   `validate:"required" json:"stock"`
	CategoryID uint   `validate:"required" json:"category_id"`
}

type ProductUpdateStockRequest struct {
	Stock uint `json:"stock" validate:"required,gte=0"`
}
