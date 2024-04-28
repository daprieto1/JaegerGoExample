package database

import (
	"context"

	"github.com/Salaton/tracing/pkg/usecase"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

type PostgresStore struct {
	db *gorm.DB
}

func NewPostgresDataStore(DB *gorm.DB) *PostgresStore {
	return &PostgresStore{
		db: DB,
	}
}

func (p PostgresStore) CreateProduct(ctx context.Context, product usecase.Product) (usecase.Product, error) {
	prod := Product{}
	err := mapstructure.Decode(product, &prod)
	if err != nil {
		return usecase.Product{}, err
	}

	err = p.db.Create(&prod).Error
	if err != nil {
		return usecase.Product{}, err
	}

	return usecase.Product{
		Name:        prod.Name,
		Description: prod.Description,
		Price:       prod.Price,
	}, nil
}
