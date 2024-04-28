package usecase

import "context"

type Storer interface {
	CreateProduct(ctx context.Context, product Product) (Product, error)
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
	return u.store.CreateProduct(ctx, product)
}
