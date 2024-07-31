package usecase

import (
	"context"

	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("github.com/daprieto1/tracing/pkg/usecases/product")

type Storer interface {
	CreateProduct(ctx context.Context, product Product) (Product, error)
	GetProductByName(ctx context.Context, name string) (Product, error)
	GetProductByDescription(ctx context.Context, description string) ([]Product, error)
}

type UseCaseImplementation struct {
	store Storer
}

func NewUseCaseImplementation(storer Storer) *UseCaseImplementation {
	return &UseCaseImplementation{
		store: storer,
	}
}

func (u UseCaseImplementation) CreateProduct(ctx context.Context, product Product) (Product, error) {
	ctx, span := tracer.Start(ctx, "CreateProduct Service")
	defer span.End()

	return u.store.CreateProduct(ctx, product)
}

func (u UseCaseImplementation) GetProductByName(ctx context.Context, name string) (Product, error) {
	ctx, span := tracer.Start(ctx, "GetProductByName Service")
	defer span.End()

	return u.store.GetProductByName(ctx, name)
}

func (u UseCaseImplementation) GetProductByDescription(ctx context.Context, description string) ([]Product, error) {
	ctx, span := tracer.Start(ctx, "GetProductByDescription Service")
	defer span.End()

	return u.store.GetProductByDescription(ctx, description)
}
