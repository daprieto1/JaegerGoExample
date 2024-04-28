package database

import "gorm.io/gorm"

type Product struct {
	gorm.Model

	Name        string  `gorm:"size:255;not null;index"`
	Description string  `gorm:"size:1024;not null"`
	Price       float64 `gorm:"not null"`
}
