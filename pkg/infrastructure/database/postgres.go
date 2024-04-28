package database

import "gorm.io/gorm"

type GenieDB struct {
	db *gorm.DB
}

func NewPostgresDataStore(DB *gorm.DB) *GenieDB {
	return &GenieDB{
		db: DB,
	}
}
