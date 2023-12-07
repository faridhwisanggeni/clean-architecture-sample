package usecase

import (
	"github.com/cleanarchitect/post/domain"
)

type TransaksiDetailRepositoryInPort interface {
	CreateTransaksiDetail(transaksi *domain.TransaksiDetail) (*domain.TransaksiDetail, error)
}

type TransaksiDetailRepositoryOutPort interface {
	CreateTransaksiDetailResponse(*domain.TransaksiDetail, error) (*domain.TransaksiDetail, error)
}
