package infrastructure

import "github.com/cleanarchitect/pos/domain"

type ProductRepository interface {
	GetProducts() ([]*domain.Product, error)
	GetProductByID(id string) (*domain.Product, error)
	CreateProduct(product *domain.Product) (*domain.Product, error)
	UpdateProduct(id string, name string, price float64) (*domain.Product, error)
	DeleteProduct(id string) error
}
