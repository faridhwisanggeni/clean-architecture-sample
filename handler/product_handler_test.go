package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cleanarchitect/pos/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// MockProductUseCase is a mock implementation of ProductRepositoryInPort for testing purposes.
type MockProductUseCase struct {
	Products   []*domain.Product
	ErrGet     error
	ErrGetByID error
	ErrCreate  error
	ErrUpdate  error
	ErrDelete  error
}

func (m *MockProductUseCase) GetProducts() ([]*domain.Product, error) {
	return m.Products, m.ErrGet
}

func (m *MockProductUseCase) GetProductByID(id string) (*domain.Product, error) {
	if m.ErrGetByID != nil {
		return nil, m.ErrGetByID
	}
	for _, product := range m.Products {
		if product.ID == id {
			return product, nil
		}
	}
	return nil, nil // Return nil when product not found
}

func (m *MockProductUseCase) CreateProduct(product *domain.Product) (*domain.Product, error) {
	return nil, m.ErrCreate
}

func (m *MockProductUseCase) UpdateProduct(id string, name string, price float64) (*domain.Product, error) {
	return nil, m.ErrUpdate
}

func (m *MockProductUseCase) DeleteProduct(id string) error {
	return m.ErrDelete
}

func TestProductHandler_GetProducts(t *testing.T) {
	// Arrange
	mockUseCase := &MockProductUseCase{}
	handler := NewProductHandler(mockUseCase)
	r := gin.Default()
	r.GET("/products", handler.GetProducts)
	req, _ := http.NewRequest("GET", "/products", nil)
	w := httptest.NewRecorder()

	// Act
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	// Add more assertions based on your expected behavior
}

func TestProductHandler_GetProductByID(t *testing.T) {
	// Arrange
	mockUseCase := &MockProductUseCase{
		Products: []*domain.Product{
			{ID: "1", Name: "Product 1", Price: 10.0},
			{ID: "2", Name: "Product 2", Price: 20.0},
		},
	}
	handler := NewProductHandler(mockUseCase)
	r := gin.Default()
	r.GET("/products/:id", handler.GetProductByID)
	req, _ := http.NewRequest("GET", "/products/1", nil)
	w := httptest.NewRecorder()

	// Act
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
	var responseProduct domain.Product
	err := json.Unmarshal(w.Body.Bytes(), &responseProduct)
	assert.Nil(t, err)
	assert.Equal(t, "1", responseProduct.ID)
	assert.Equal(t, "Product 1", responseProduct.Name)
	assert.Equal(t, 10.0, responseProduct.Price)
}

func TestProductHandler_CreateProduct(t *testing.T) {
	// Arrange
	mockUseCase := &MockProductUseCase{}
	handler := NewProductHandler(mockUseCase)
	r := gin.Default()
	r.POST("/products", handler.CreateProduct)
	reqBody := []byte(`{"name":"New Product","price":15.0}`)
	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestProductHandler_UpdateProduct(t *testing.T) {
	// Arrange
	mockUseCase := &MockProductUseCase{}
	handler := NewProductHandler(mockUseCase)
	r := gin.Default()
	r.PUT("/products/:id", handler.UpdateProduct)
	reqBody := []byte(`{"name":"Updated Product","price":20.0}`)
	req, _ := http.NewRequest("PUT", "/products/123", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestProductHandler_DeleteProduct(t *testing.T) {
	// Arrange
	mockUseCase := &MockProductUseCase{}
	handler := NewProductHandler(mockUseCase)
	r := gin.Default()
	r.DELETE("/products/:id", handler.DeleteProduct)
	req, _ := http.NewRequest("DELETE", "/products/123", nil)
	w := httptest.NewRecorder()

	// Act
	r.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusNoContent, w.Code)
}
