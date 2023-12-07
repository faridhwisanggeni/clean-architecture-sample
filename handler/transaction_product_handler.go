package handler

import (
	"fmt"
	"net/http"

	"github.com/cleanarchitect/pos/domain"
	"github.com/cleanarchitect/pos/usecase"
	"github.com/gin-gonic/gin"
)

type TransactionProductHandler struct {
	TransactionProductUseCase usecase.TransactionRepositoryInPort
}

func NewTransactionHandler(transactionUseCase usecase.TransactionRepositoryInPort) *TransactionProductHandler {
	return &TransactionProductHandler{
		TransactionProductUseCase: transactionUseCase,
	}
}

//controller
//implement CRUD
//get transactions handles the request
func (h *TransactionProductHandler) GetTransaction(c *gin.Context) {
	trx, err := h.TransactionProductUseCase.GetTransaction()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "There are no transactions"})
		return
	}
	c.JSON(http.StatusOK, trx)
}

//get transaction by id handles the request by id
func (h *TransactionProductHandler) GetTransactionByID(c *gin.Context) {
	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
	// 	return
	// }
	trx, err := h.TransactionProductUseCase.GetTransactionByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, trx)
}

//create product handles request add new product
func (h *TransactionProductHandler) CreateTransaction(c *gin.Context) {
	var trx domain.TransactionProduct
	if err := c.ShouldBindJSON(&trx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body request"})
		return
	}
	fmt.Println(trx)
	createTrx, err := h.TransactionProductUseCase.CreateTransaction(&trx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving transaction, please try again"})
		return
	}

	c.JSON(http.StatusCreated, createTrx)
}

//update product handle request update product
// func (h *ProductHandler) UpdateProduct(c *gin.Context) {
// 	var updatedProduct domain.Product
// 	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body request"})
// 		return
// 	}

// 	updateProduct, err := h.ProductUseCase.UpdateProduct(&updatedProduct)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "update product failed, please try again"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, updateProduct)
// }

// //delete product handles request for delete product
// func (h *ProductHandler) DeleteProduct(c *gin.Context) {
// 	err := h.ProductUseCase.DeleteProduct(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete product failed, please try again"})
// 		return
// 	}

// 	c.JSON(http.StatusNoContent, nil)
// }
