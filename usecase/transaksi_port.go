package usecase

import (
	"github.com/cleanarchitect/post/domain"
)

type TransaksiRepositoryInPort interface {
	CreateTransaksi(transaksi *domain.Transaksi) (*domain.Transaksi, error)
}

type TransaksiRepositoryOutPort interface {
	CreateTransaksiResponse(*domain.Transaksi, error) (*domain.Transaksi, error)
}
