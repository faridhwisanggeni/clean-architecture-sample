package usecase

import (
	"fmt"
	"github.com/cleanarchitect/post/domain"
	"strconv"
)

type OrderUseCaseInteractor struct {
	OrderRepositoryInPort    OrderRepositoryInPort
	OrderRepositoryOutPort   OrderRepositoryOutPort
	ProductRepositoryInPort  ProductRepositoryInPort
	ProductRepositoryOutPort ProductRepositoryOutPort
	TrxRepositoryInPort      TrxRepositoryInPort
	TrxRepositoryOutPort     TrxRepositoryOutPort
}

func NewOrderUseCase(orderRepositoryInPort OrderRepositoryInPort, orderRepositoryOutPort OrderRepositoryOutPort, productRepositoryInPort ProductRepositoryInPort, productRepositoryOutPort ProductRepositoryOutPort) *OrderUseCaseInteractor {
	return &OrderUseCaseInteractor{OrderRepositoryInPort: orderRepositoryInPort, OrderRepositoryOutPort: orderRepositoryOutPort, ProductRepositoryInPort: productRepositoryInPort, ProductRepositoryOutPort: productRepositoryOutPort}
}

func (o *OrderUseCaseInteractor) CreateOrder(order *domain.Order) (*domain.Order, error) {
	var totalPrice float64

	for _, item := range order.Orders {
		product, err := o.ProductRepositoryOutPort.GetProductByIDResponse(o.ProductRepositoryInPort.GetProductByID(item.ID))
		if err != nil {
			return o.OrderRepositoryOutPort.CraeteOrderResponse(nil, err)
		}

		totalPrice = product.Price * float64(item.Quantity)
	}

	order.TotalPrice = totalPrice

	fmt.Println(order.TotalPrice)

	orderRes, err := o.OrderRepositoryOutPort.CraeteOrderResponse(o.OrderRepositoryInPort.CreateOrder(order))
	if err != nil {
		return o.OrderRepositoryOutPort.CraeteOrderResponse(nil, err)
	}

	for _, item := range order.Orders {
		_, err := o.TrxRepositoryOutPort.CreateTrxResponse(o.TrxRepositoryInPort.CreateTrx(&domain.Transaksi{
			ProductId:  item.ID,
			Quantity:   item.Quantity,
			OrderId:    strconv.Itoa(order.Id),
			TotalPrice: item.Price * float64(item.Quantity),
		}))
		if err != nil {
			return o.OrderRepositoryOutPort.CraeteOrderResponse(nil, err)
		}
	}

	return orderRes, nil
}
