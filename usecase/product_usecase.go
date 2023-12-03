package usecase

import (
	"github.com/cleanarchitect/pos/domain"
	"github.com/cleanarchitect/pos/infrastructure"
)

type ProductService struct {
	productRepository infrastructure.ProductRepository
}

func NewProductService(productRepository infrastructure.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

// Implement the CRUD methods

// GetProducts retrieves all products
func (s *ProductService) GetProducts() ([]*domain.Product, error) {
	return s.productRepository.GetProducts()
}

// GetProductByID retrieves a product by ID
func (s *ProductService) GetProductByID(id string) (*domain.Product, error) {
	return s.productRepository.GetProductByID(id)
}

// CreateProduct inserts a new product into the database
func (s *ProductService) CreateProduct(product *domain.Product) (*domain.Product, error) {
	return s.productRepository.CreateProduct(product)
}

// UpdateProduct updates a product in the database
func (s *ProductService) UpdateProduct(id string, name string, price float64) (*domain.Product, error) {
	return s.productRepository.UpdateProduct(id, name, price)
}

// DeleteProduct deletes a product from the database
func (s *ProductService) DeleteProduct(id string) error {
	return s.productRepository.DeleteProduct(id)
}
