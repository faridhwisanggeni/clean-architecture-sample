package handler

import (
	"net/http"

	"github.com/cleanarchitect/pos/domain"
	"github.com/cleanarchitect/pos/usecase"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUseCase usecase.ProductRepositoryInPort
}

func NewProductHandler(productUseCase usecase.ProductRepositoryInPort) *ProductHandler {
	return &ProductHandler{
		productUseCase: productUseCase,
	}
}

// Controllers
// Implement the CRUD handlers
// GetProducts handles the request to get all products
func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.productUseCase.GetProducts() //product is presenters from outport
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetProductByID handles the request to get a product by ID
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	// id := strconv.Itoa(c.Param("id"))
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
	// 	return
	// }

	product, err := h.productUseCase.GetProductByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// CreateProduct handles the request to create a new product
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdProduct, err := h.productUseCase.CreateProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}

// UpdateProduct handles the request to update a product
func (h *ProductHandler) UpdateProduct(c *gin.Context) {

	var updatedProduct domain.Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateProduct, err := h.productUseCase.UpdateProduct(updatedProduct.ID, updatedProduct.Name, updatedProduct.Price)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, updateProduct)
}

// DeleteProduct handles the request to delete a product
func (h *ProductHandler) DeleteProduct(c *gin.Context) {

	err := h.productUseCase.DeleteProduct(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
