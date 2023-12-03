package usecase

import (
	"github.com/cleanarchitect/pos/domain"
	"github.com/cleanarchitect/pos/infrastructure"
)

type ProductUseCaseImpl struct {
	productRepository infrastructure.ProductRepository
}

func NewProductUseCase(productRepository infrastructure.ProductRepository) *ProductUseCaseImpl {
	return &ProductUseCaseImpl{
		productRepository: productRepository,
	}
}

func (uc *ProductUseCaseImpl) GetProducts() ([]*domain.Product, error) {
	return uc.productRepository.GetProducts()
}

func (uc *ProductUseCaseImpl) GetProductByID(id string) (*domain.Product, error) {
	return uc.productRepository.GetProductByID(id)
}

func (uc *ProductUseCaseImpl) CreateProduct(product *domain.Product) (*domain.Product, error) {
	return uc.productRepository.CreateProduct(product)
}

func (uc *ProductUseCaseImpl) UpdateProduct(id string, name string, price float64) (*domain.Product, error) {
	return uc.productRepository.UpdateProduct(id, name, price)
}

func (uc *ProductUseCaseImpl) DeleteProduct(id string) error {
	return uc.productRepository.DeleteProduct(id)
}
