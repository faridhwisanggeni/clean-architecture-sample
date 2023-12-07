package handler

import (
	"net/http"

	"github.com/cleanarchitect/pos/domain"
	"github.com/cleanarchitect/pos/usecase"
	"github.com/gin-gonic/gin"
)

type TransactionDetailHandler struct {
	TransactionDetailUseCase usecase.TransactionDetailRepositoryInPort
}

func NewTransactionDetailHandler(transactionUseCase usecase.TransactionDetailRepositoryInPort) *TransactionDetailHandler {
	return &TransactionDetailHandler{
		TransactionDetailUseCase: transactionUseCase,
	}
}

//controller
//implement CRUD
//get transactions handles the request
func (h *TransactionDetailHandler) GetTransactionDetail(c *gin.Context) {
	trx, err := h.TransactionDetailUseCase.GetTransactionDetail()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "There are no transactions detail"})
		return
	}
	c.JSON(http.StatusOK, trx)
}

//get transaction by id handles the request by id
func (h *TransactionDetailHandler) GetTransactionDetailByTrxID(c *gin.Context) {
	// id, err := strconv.Atoi(c.Param("id"))
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid trx id"})
	// 	return
	// }
	trx, err := h.TransactionDetailUseCase.GetTransactionDetailByTrxID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction detail not found"})
		return
	}

	c.JSON(http.StatusOK, trx)
}

//create product handles request add new product
func (h *TransactionDetailHandler) CreateTransactionDetail(c *gin.Context) {
	var trx domain.TransactionDetail
	if err := c.ShouldBindJSON(&trx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body request"})
		return
	}

	createTrx, err := h.TransactionDetailUseCase.CreateTransactionDetail(&trx)
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
