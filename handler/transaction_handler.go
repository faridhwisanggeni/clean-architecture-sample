package handler

import (
	"net/http"

	"github.com/cleanarchitect/pos/domain"
	"github.com/cleanarchitect/pos/usecase"
	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	transactionUseCase usecase.TransactionRepositoryInPort
}

func NewTransactionHandler(transactionUseCase usecase.TransactionRepositoryInPort) *TransactionHandler {
	return &TransactionHandler{
		transactionUseCase: transactionUseCase,
	}
}

// Controllers
// Create incoming transaction order request
func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	var transaction domain.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTrx, err := h.transactionUseCase.CreateTransaction(&transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdTrx)
}

func (h *TransactionHandler) CreateTransactionDetail(c *gin.Context) {
	var transaction domain.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTrx, err := h.transactionUseCase.CreateTransaction(&transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdTrx)
}
