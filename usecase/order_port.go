package usecase

import "github.com/cleanarchitect/post/domain"

type OrderRepositoryInPort interface {
	CreateOrder(order *domain.Order) (*domain.Order, error)
}

type OrderRepositoryOutPort interface {
	CraeteOrderResponse(data *domain.Order, err error) (*domain.Order, error)
}
