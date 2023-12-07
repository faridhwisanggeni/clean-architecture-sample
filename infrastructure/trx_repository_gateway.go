package infrastructure

import (
	"context"
	"fmt"
	"github.com/cleanarchitect/post/domain"
)

type TrxRepositoryImpl struct {
	dbConn *pgx.Conn
}

func NewTrxRepository(dbConn *pgx.Conn) *TrxRepositoryImpl {
	return &TrxRepositoryImpl{dbConn: dbConn}
}

func (r *TrxRepositoryImpl) CreateTrx(transaksi *domain.Transaksi) (*domain.Transaksi, error) {
	sql := "insert into trx(product_id, quantity, order_id, totalprice) VALUES($1,$2,$3,$4) returning id"

	row := r.dbConn.QueryRow(context.Background(), sql, transaksi.ProductId, transaksi.Quantity, transaksi.OrderId, transaksi.TotalPrice)

	var createdTrx domain.Transaksi
	if err := row.Scan(&createdTrx.Id); err != nil {
		return nil, fmt.Errorf("error scanning transaksi: %w", err)
	}

	return &createdTrx, nil
}

func (r *TrxRepositoryImpl) CraeteOrderResponse(data *domain.Order, err error) (*domain.Order, error) {
	return data, nil
}
