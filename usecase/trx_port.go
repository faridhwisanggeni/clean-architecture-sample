package usecase

import "github.com/cleanarchitect/post/domain"

type TrxRepositoryInPort interface {
	CreateTrx(transaksi *domain.Transaksi) (*domain.Transaksi, error)
}

type TrxRepositoryOutPort interface {
	CreateTrxResponse(transaksi *domain.Transaksi, err error) (*domain.Transaksi, error)
}
