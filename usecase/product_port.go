package usecase

import "github.com/cleanarchitect/pos/domain"

type ProductRepositoryInPort interface {
	GetProducts() ([]*domain.Product, error)
	GetProductByID(id string) (*domain.Product, error)
	CreateProduct(product *domain.Product) (*domain.Product, error)
	UpdateProduct(id string, name string, price float64) (*domain.Product, error)
	DeleteProduct(id string) error
}

type ProductRepositoryOutPort interface {
	GetProductsResponse([]*domain.Product, error) ([]*domain.Product, error)
	GetProductByIDResponse(*domain.Product, error) (*domain.Product, error)
	CreateProductResponse(*domain.Product, error) (*domain.Product, error)
	UpdateProductResponse(*domain.Product, error) (*domain.Product, error)
	DeleteProductResponse(error) error
}
