package database

import (
	"context"

	"github.com/daprieto1/tracing/pkg/usecase"
	"github.com/mitchellh/mapstructure"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

var tracer = otel.Tracer("github.com/daprieto1/tracing/pkg/infrastructure/database/postgres")

type PostgresStore struct {
	db *gorm.DB
}

func NewPostgresDataStore(DB *gorm.DB) *PostgresStore {
	return &PostgresStore{
		db: DB,
	}
}

func (p PostgresStore) CreateProduct(ctx context.Context, product usecase.Product) (usecase.Product, error) {
	_, span := tracer.Start(ctx, "CreateProduct Database")
	defer span.End()

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

func (p PostgresStore) GetProductByName(ctx context.Context, name string) (usecase.Product, error) {
	_, span := tracer.Start(ctx, "GetProductByName Database")
	defer span.End()

	prod := Product{}
	result := p.db.Where("name = ?", name).First(&prod)
	if result.Error != nil {
		return usecase.Product{}, result.Error
	}

	return usecase.Product{
		Name:        prod.Name,
		Description: prod.Description,
		Price:       prod.Price,
	}, nil
}

func (p PostgresStore) GetProductByDescription(ctx context.Context, description string) ([]usecase.Product, error) {
	_, span := tracer.Start(ctx, "GetProductByDescription Database")
	defer span.End()

	var prods []Product
	result := p.db.Where("description LIKE ?", "%"+description+"%").Order("name ASC").Find(&prods)
	if result.Error != nil {
		return []usecase.Product{}, result.Error
	}

	results := []usecase.Product{}
	for _, prod := range prods {
		results = append(results, usecase.Product{
			Name:        prod.Name,
			Description: prod.Description,
			Price:       prod.Price,
		})
	}
	return results, nil

}
