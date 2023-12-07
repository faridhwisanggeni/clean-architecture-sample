package infrastructure

import (
	"context"
	"fmt"

	"github.com/cleanarchitect/post/domain"
	"github.com/jackc/pgx/v4"
)

type TransaksiRepositoryImpl struct {
	dbConn *pgx.Conn
}

func NewTransaksiRepository(dbConn *pgx.Conn) *TransaksiRepositoryImpl {
	return &TransaksiRepositoryImpl{dbConn: dbConn}
}

func (r *TransaksiRepositoryImpl) CreateTransaksi(transaksi *domain.Transaksi) (*domain.Transaksi, error) {
	sql := "insert into transaksi(total_qty,total_price) values ($1,$2) returning id,total_qty,total_price"

	row := r.dbConn.QueryRow(context.Background(), sql, transaksi.TotalQty, transaksi.TotalPrice)
	var createdTransaksi domain.Transaksi
	if err := row.Scan(&createdTransaksi.ID, &createdTransaksi.TotalQty, &createdTransaksi.TotalPrice); err != nil {
		return nil, fmt.Errorf("error scanning transaksi: %w", err)
	}
	return &createdTransaksi, nil
}

func (r *TransaksiRepositoryImpl) CreateTransaksiResponse(data *domain.Transaksi, err error) (*domain.Transaksi, error) {
	return data, err
}
