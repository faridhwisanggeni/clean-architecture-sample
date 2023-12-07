package handler

import (
	"fmt"
	"net/http"

	"github.com/cleanarchitect/post/domain"
	"github.com/cleanarchitect/post/usecase"
	"github.com/gin-gonic/gin"
)

type TransaksiHandler struct {
	transaksiUseCase usecase.TransaksiRepositoryInPort
}

func NewTransaksitHandler(transaksiUseCase usecase.TransaksiRepositoryInPort) *TransaksiHandler {
	return &TransaksiHandler{
		transaksiUseCase: transaksiUseCase,
	}
}

func (h TransaksiHandler) CreateTransaksi(c *gin.Context) {
	var transaksi domain.Transaksi

	if err := c.ShouldBindJSON(&transaksi); err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	println(&transaksi)
	fmt.Println(transaksi)
	fmt.Println(c.Request.Body)

	totalPrice := 0
	totalQty := 0
	var trxDet usecase.TransaksiDetailRepositoryInPort
	valid := true
	for _, key := range transaksi.Item {
		fmt.Println("key ", key)

		totalPrice = totalPrice + (int(key.BasicPrice) * key.Qty)
		totalQty = totalQty + key.Qty

		var det domain.TransaksiDetail
		det.BasicPrice = key.BasicPrice
		det.IDProduct = key.IDProduct
		det.Qty = key.Qty
		det.Total = key.BasicPrice * float64(key.Qty)

		_, err := trxDet.CreateTransaksiDetail(&det)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	fmt.Println("valid", valid)

	var trx domain.Transaksi
	trx.TotalPrice = float64(totalPrice)
	trx.TotalQty = totalQty

	trxh, err := h.transaksiUseCase.CreateTransaksi(&trx)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, trxh)
}
