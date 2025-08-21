package domain

import "gorm.io/gorm"

type ProductCategory struct {
	ID        uint `gorm:"primaryKey"`
	Category  string
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
