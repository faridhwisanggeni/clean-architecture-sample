package handler

import (
	"github.com/cleanarchitect/post/domain"
	"github.com/cleanarchitect/post/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderHandler struct {
	orderUseCase usecase.OrderRepositoryInPort
}

func NewOrderHandler(orderUseCase usecase.OrderRepositoryInPort) *OrderHandler {
	return &OrderHandler{orderUseCase: orderUseCase}
}

func (o *OrderHandler) CreateProduct(c *gin.Context) {
	var order domain.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//createdProduct, err := o.productUseCase.CreateProduct(&product)
	createOrder, err := o.orderUseCase.CreateOrder(&order)

	if err != nil {
		//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createOrder)
}
