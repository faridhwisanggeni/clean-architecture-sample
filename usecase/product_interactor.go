package usecase

import (
	"github.com/cleanarchitect/pos/domain"
)

type ProductInteractor struct {
	productRepository ProductRepositoryInPort
}

func NewProductInteractor(productRepository ProductRepositoryInPort) *ProductInteractor {
	return &ProductInteractor{
		productRepository: productRepository,
	}
}

// Implement the CRUD methods contract

// GetProducts retrieves all products
func (s *ProductInteractor) GetProducts() ([]*domain.Product, error) {
	return s.productRepository.GetProducts()
}

// GetProductByID retrieves a product by ID
func (s *ProductInteractor) GetProductByID(id string) (*domain.Product, error) {
	return s.productRepository.GetProductByID(id)
}

// CreateProduct inserts a new product into the database
func (s *ProductInteractor) CreateProduct(product *domain.Product) (*domain.Product, error) {
	return s.productRepository.CreateProduct(product)
}

// UpdateProduct updates a product in the database
func (s *ProductInteractor) UpdateProduct(id string, name string, price float64) (*domain.Product, error) {
	return s.productRepository.UpdateProduct(id, name, price)
}

// DeleteProduct deletes a product from the database
func (s *ProductInteractor) DeleteProduct(id string) error {
	return s.productRepository.DeleteProduct(id)
}
