package usecase

import (
	"fmt"

	"github.com/cleanarchitect/pos/domain"
)

type ProductUseCaseImpl struct {
	productRepositoryIn  ProductRepositoryInPort
	productRepositoryOut ProductRepositoryOutPort
}

func NewProductUseCase(productRepositoryIn ProductRepositoryInPort, productRepositoryOut ProductRepositoryOutPort) *ProductUseCaseImpl {
	return &ProductUseCaseImpl{
		productRepositoryIn:  productRepositoryIn,
		productRepositoryOut: productRepositoryOut,
	}
}

func (uc *ProductUseCaseImpl) GetProducts() ([]*domain.Product, error) {
	return uc.productRepositoryOut.GetProductsResponse(uc.productRepositoryIn.GetProducts())
}

func (uc *ProductUseCaseImpl) GetProductByID(id string) (*domain.Product, error) {
	return uc.productRepositoryOut.GetProductByIDResponse(uc.productRepositoryIn.GetProductByID(id))
}

func (uc *ProductUseCaseImpl) CreateProduct(product *domain.Product) (*domain.Product, error) {
	if product.Price > 200 {
		var product = &domain.Product{}
		return uc.productRepositoryOut.CreateProductResponse(product, fmt.Errorf("Nilai price kemahalan"))
	}
	return uc.productRepositoryOut.CreateProductResponse(uc.productRepositoryIn.CreateProduct(product))
}

func (uc *ProductUseCaseImpl) UpdateProduct(id string, name string, price float64) (*domain.Product, error) {
	return uc.productRepositoryOut.UpdateProductResponse(uc.productRepositoryIn.UpdateProduct(id, name, price))
}

func (uc *ProductUseCaseImpl) DeleteProduct(id string) error {
	return uc.productRepositoryOut.DeleteProductResponse(uc.productRepositoryIn.DeleteProduct(id))
}

func (uc *ProductUseCaseImpl) GetProductsResponse() ([]*domain.Product, error) {
	return uc.productRepositoryOut.GetProductsResponse(uc.productRepositoryIn.GetProducts())
}

func (uc *ProductUseCaseImpl) GetProductByIDResponse(id string) (*domain.Product, error) {
	return uc.productRepositoryOut.GetProductByIDResponse(uc.productRepositoryIn.GetProductByID(id))
}

func (uc *ProductUseCaseImpl) CreateProductResponse(product *domain.Product) (*domain.Product, error) {
	return uc.productRepositoryOut.CreateProductResponse(uc.productRepositoryIn.CreateProduct(product))
}

func (uc *ProductUseCaseImpl) UpdateProductResponse(id string, name string, price float64) (*domain.Product, error) {
	return uc.productRepositoryOut.UpdateProductResponse(uc.productRepositoryIn.UpdateProduct(id, name, price))
}

func (uc *ProductUseCaseImpl) DeleteProductResponse(id string) error {
	return uc.productRepositoryOut.DeleteProductResponse(uc.productRepositoryIn.DeleteProduct(id))
}
