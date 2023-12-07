package usecase

import (
	"fmt"

	"github.com/cleanarchitect/pos/domain"
)

type ProductUseCaseInteractor struct {
	productRepositoryIn  ProductRepositoryInPort
	productRepositoryOut ProductRepositoryOutPort
}

func NewProductUseCase(productRepositoryIn ProductRepositoryInPort, productRepositoryOut ProductRepositoryOutPort) *ProductUseCaseInteractor {
	return &ProductUseCaseInteractor{
		productRepositoryIn:  productRepositoryIn,
		productRepositoryOut: productRepositoryOut,
	}
}

func (uc *ProductUseCaseInteractor) GetProducts() ([]*domain.Product, error) {
	return uc.productRepositoryOut.GetProductsResponse(uc.productRepositoryIn.GetProducts())
}

func (uc *ProductUseCaseInteractor) GetProductByID(id string) (*domain.Product, error) {
	return uc.productRepositoryOut.GetProductByIDResponse(uc.productRepositoryIn.GetProductByID(id))
}

func (uc *ProductUseCaseInteractor) CreateProduct(product *domain.Product) (*domain.Product, error) {
	if product.Price > 200 {
		var product = &domain.Product{}
		return uc.productRepositoryOut.CreateProductResponse(product, fmt.Errorf("Nilai price kemahalan"))
	}
	if product.Quantity > 10 {
		var product = &domain.Product{}
		return uc.productRepositoryOut.CreateProductResponse(product, fmt.Errorf("Quantity kebanyakan, mobilnya gak muat"))
	}
	return uc.productRepositoryOut.CreateProductResponse(uc.productRepositoryIn.CreateProduct(product))
}

func (uc *ProductUseCaseInteractor) UpdateProduct(id string, name string, price float64) (*domain.Product, error) {
	return uc.productRepositoryOut.UpdateProductResponse(uc.productRepositoryIn.UpdateProduct(id, name, price))
}

func (uc *ProductUseCaseInteractor) DeleteProduct(id string) error {
	return uc.productRepositoryOut.DeleteProductResponse(uc.productRepositoryIn.DeleteProduct(id))
}

func (uc *ProductUseCaseInteractor) GetProductsResponse() ([]*domain.Product, error) {
	return uc.productRepositoryOut.GetProductsResponse(uc.productRepositoryIn.GetProducts())
}

func (uc *ProductUseCaseInteractor) GetProductByIDResponse(id string) (*domain.Product, error) {
	return uc.productRepositoryOut.GetProductByIDResponse(uc.productRepositoryIn.GetProductByID(id))
}

func (uc *ProductUseCaseInteractor) CreateProductResponse(product *domain.Product, err error) (*domain.Product, error) {
	return uc.productRepositoryOut.CreateProductResponse(uc.productRepositoryIn.CreateProduct(product))
}

func (uc *ProductUseCaseInteractor) UpdateProductResponse(id string, name string, price float64) (*domain.Product, error) {
	return uc.productRepositoryOut.UpdateProductResponse(uc.productRepositoryIn.UpdateProduct(id, name, price))
}

func (uc *ProductUseCaseInteractor) DeleteProductResponse(id string) error {
	return uc.productRepositoryOut.DeleteProductResponse(uc.productRepositoryIn.DeleteProduct(id))
}
