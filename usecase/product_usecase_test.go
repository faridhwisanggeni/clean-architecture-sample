package usecase

import (
	"fmt"
	"testing"

	"github.com/cleanarchitect/pos/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockProductRepository is a mock implementation of ProductRepositoryInPort and ProductRepositoryOutPort for testing purposes.
type MockProductRepository struct {
	Products   []*domain.Product
	ErrGet     error
	ErrGetByID error
	ErrCreate  error
	ErrUpdate  error
	ErrDelete  error
}

// Mocking Function
type MockProductRepositoryInPort struct {
	mock.Mock
}

func (m *MockProductRepositoryInPort) GetProducts() ([]*domain.Product, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Product), args.Error(1)
}

func (m *MockProductRepositoryInPort) GetProductByID(id string) (*domain.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *MockProductRepositoryInPort) CreateProduct(product *domain.Product) (*domain.Product, error) {
	args := m.Called(product)

	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *MockProductRepositoryInPort) UpdateProduct(id string, name string, price float64) (*domain.Product, error) {
	args := m.Called(id, name, price)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *MockProductRepositoryInPort) DeleteProduct(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

type MockProductRepositoryOutPort struct {
	mock.Mock
}

func (m *MockProductRepositoryOutPort) GetProductsResponse(products []*domain.Product, err error) ([]*domain.Product, error) {
	args := m.Called(products, err)
	return args.Get(0).([]*domain.Product), args.Error(1)
}

func (m *MockProductRepositoryOutPort) GetProductByIDResponse(product *domain.Product, err error) (*domain.Product, error) {
	args := m.Called(product, err)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *MockProductRepositoryOutPort) CreateProductResponse(product *domain.Product, err error) (*domain.Product, error) {
	args := m.Called(product, err)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *MockProductRepositoryOutPort) UpdateProductResponse(product *domain.Product, err error) (*domain.Product, error) {
	args := m.Called(product, err)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *MockProductRepositoryOutPort) DeleteProductResponse(err error) error {
	args := m.Called(err)
	return args.Error(0)
}

func TestProductUseCase_GetProducts(t *testing.T) {
	// Arrange
	mockRepositoryIn := &MockProductRepositoryInPort{}
	mockRepositoryOut := &MockProductRepositoryOutPort{}
	useCase := NewProductUseCase(mockRepositoryIn, mockRepositoryOut)

	expectedProducts := []*domain.Product{
		{ID: "1", Name: "Product1", Price: 20.5},
		{ID: "2", Name: "Product2", Price: 30.75},
	}

	// expectedResults, err := mockRepositoryOut.GetProductsResponse(expectedProducts, nil)
	mockRepositoryOut.On("GetProductsResponse", expectedProducts, nil).Return(expectedProducts, nil)
	// Use GetProducts from MockProductRepositoryOutPort
	mockRepositoryIn.On("GetProducts").Return(expectedProducts, nil)

	// Act
	products, err := useCase.GetProducts()

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, products)
	assert.Equal(t, expectedProducts, products)

	// Verify that the expected method was called
	mockRepositoryOut.AssertExpectations(t)
}

func TestProductUseCase_GetProductByID(t *testing.T) {
	// Arrange
	mockRepositoryIn := &MockProductRepositoryInPort{}
	mockRepositoryOut := &MockProductRepositoryOutPort{}
	useCase := NewProductUseCase(mockRepositoryIn, mockRepositoryOut)

	expectedProduct := &domain.Product{ID: "1", Name: "Product1", Price: 20.5}

	// Mocking the response of GetProductByID from MockProductRepositoryInPort
	mockRepositoryIn.On("GetProductByID", "1").Return(expectedProduct, nil)

	// Mocking the response of GetProductByIDResponse from MockProductRepositoryOutPort
	mockRepositoryOut.On("GetProductByIDResponse", expectedProduct, nil).
		Return(expectedProduct, nil)

	// Act
	product, err := useCase.GetProductByIDResponse("1")

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, expectedProduct, product)

	// Verify that the expected methods were called
	mockRepositoryIn.AssertExpectations(t)
	mockRepositoryOut.AssertExpectations(t)
}

func TestProductUseCase_GetProductByID_Error(t *testing.T) {
	// Arrange
	mockRepositoryIn := &MockProductRepositoryInPort{}
	mockRepositoryOut := &MockProductRepositoryOutPort{}
	useCase := NewProductUseCase(mockRepositoryIn, mockRepositoryOut)

	expectedError := fmt.Errorf("not found")
	expectedProducts := &domain.Product{}

	// Mocking the response of GetProductByID from MockProductRepositoryInPort to return an error
	mockRepositoryIn.On("GetProductByID", "1").Return(expectedProducts, expectedError)

	// Mocking the response of GetProductByIDResponse from MockProductRepositoryOutPort to return an error
	mockRepositoryOut.On("GetProductByIDResponse", expectedProducts, expectedError).
		Return(expectedProducts, expectedError)

	// Act
	product, err := useCase.GetProductByID("1")

	// Assert
	assert.Error(t, err)
	assert.NotNil(t, product)
	// assert.Equal(t, expectedError, err)

	// Verify that the expected methods were called
	mockRepositoryIn.AssertExpectations(t)
	mockRepositoryOut.AssertExpectations(t)
}

func TestProductUseCase_CreateProduct(t *testing.T) {
	// Arrange
	mockRepositoryIn := &MockProductRepositoryInPort{}
	mockRepositoryOut := &MockProductRepositoryOutPort{}
	useCase := NewProductUseCase(mockRepositoryIn, mockRepositoryOut)

	// Positive scenario
	product := &domain.Product{ID: "1", Name: "TestProduct", Price: 100, Quantity: 5}
	mockRepositoryIn.On("CreateProduct", product).Return(product, nil)
	mockRepositoryOut.On("CreateProductResponse", product, nil).Return(product, nil)
	// Act
	createdProduct, err := useCase.CreateProduct(product)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, createdProduct)
	assert.Equal(t, "1", createdProduct.ID)
}

func TestProductUseCase_CreateProduct_Error(t *testing.T) {
	// Arrange
	mockRepositoryIn := &MockProductRepositoryInPort{}
	mockRepositoryOut := &MockProductRepositoryOutPort{}
	useCase := NewProductUseCase(mockRepositoryIn, mockRepositoryOut)

	// Negative scenario with high price
	product := &domain.Product{ID: "1", Name: "TestProduct", Price: 300, Quantity: 5}
	productEmpty := &domain.Product{ID: "", Name: "", Price: 0, Quantity: 0}
	expectedError := fmt.Errorf("Nilai price kemahalan")
	mockRepositoryIn.On("CreateProduct", product).Return(productEmpty, expectedError)
	mockRepositoryOut.On("CreateProductResponse", productEmpty, expectedError).Return(productEmpty, expectedError)

	// Act
	createdProduct, err := useCase.CreateProduct(product)

	// Assert
	assert.Error(t, err)
	assert.NotNil(t, createdProduct)
	assert.EqualError(t, err, expectedError.Error())
}

func TestProductUseCase_CreateProduct_Error_Quantity(t *testing.T) {
	// Arrange
	mockRepositoryIn := &MockProductRepositoryInPort{}
	mockRepositoryOut := &MockProductRepositoryOutPort{}
	useCase := NewProductUseCase(mockRepositoryIn, mockRepositoryOut)

	// Negative scenario with high quantity
	product := &domain.Product{Price: 100, Quantity: 15}
	expectedError := fmt.Errorf("Quantity kebanyakan, mobilnya gak muat")
	productEmpty := &domain.Product{ID: "", Name: "", Price: 0, Quantity: 0}
	mockRepositoryIn.On("CreateProduct", product).Return(productEmpty, expectedError)
	mockRepositoryOut.On("CreateProductResponse", productEmpty, expectedError).Return(productEmpty, expectedError)
	// Act
	createdProduct, err := useCase.CreateProduct(product)

	// Assert
	assert.Error(t, err)
	assert.Empty(t, createdProduct)
	assert.EqualError(t, err, expectedError.Error())
}
