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
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
