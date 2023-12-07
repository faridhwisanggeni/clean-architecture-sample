package infrastructure

import (
	"context"
	"fmt"
	"github.com/cleanarchitect/post/domain"
)

type OrderRepositoryImpl struct {
	dbConn *pgx.Conn
}

func NewOrderRepository(dbConn *pgx.Conn) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{dbConn: dbConn}
}

func (o *OrderRepositoryImpl) CreateOrder(order *domain.Order) (*domain.Order, error) {
	sql := "insert into orders(id,total_price) values ($1,$2) returning id"

	row := o.dbConn.QueryRow(context.Background(), sql, order.Id, order.TotalPrice)
	var createdProduct domain.Order
	if err := row.Scan(&createdProduct.Id); err != nil {
		return nil, fmt.Errorf("error scanning order: %w", err)
	}

	return &createdProduct, nil
}

func (o *OrderRepositoryImpl) CraeteOrderResponse(data *domain.Order, err error) (*domain.Order, error) {
	return data, err
}
