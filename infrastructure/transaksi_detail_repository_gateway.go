package infrastructure

import (
	"context"
	"fmt"

	"github.com/cleanarchitect/post/domain"
	"github.com/jackc/pgx/v4"
)

type TransaksiDetailRepositoryImpl struct {
	dbConn *pgx.Conn
}

func NewTransaksiDetailRepository(dbConn *pgx.Conn) *TransaksiDetailRepositoryImpl {
	return &TransaksiDetailRepositoryImpl{dbConn: dbConn}
}

func (r *TransaksiDetailRepositoryImpl) CreateTransaksi(transaksiDetails *domain.TransaksiDetail) (*domain.TransaksiDetail, error) {
	sql := "insert into transaksi_detail(id_transaksi,id_product,qty,basic_price,total) values ($1,$2,$3,$4,$5) returning id,id_transaksi,id_product,qty,basic_price,total"

	row := r.dbConn.QueryRow(context.Background(), sql, transaksiDetails.IDTransaksi, transaksiDetails.IDProduct, transaksiDetails.Qty, transaksiDetails.Total)
	var createdTransaksiDetails domain.TransaksiDetail
	if err := row.Scan(&transaksiDetails.ID, &transaksiDetails.IDTransaksi, &transaksiDetails.IDProduct, &transaksiDetails.Qty, &transaksiDetails.BasicPrice, &transaksiDetails.Total); err != nil {
		return nil, fmt.Errorf("error scanning transaksi: %w", err)
	}
	return &createdTransaksiDetails, nil
}

func (r *TransaksiDetailRepositoryImpl) CreateTransaksiResponse(data *domain.TransaksiDetail, err error) (*domain.TransaksiDetail, error) {
	return data, err
}
