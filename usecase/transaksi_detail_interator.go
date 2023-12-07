package usecase

import (
	"github.com/cleanarchitect/post/domain"
)

type TransaksiDetailUseCaseInteractor struct {
	transaksiDetailRepositoryIn  TransaksiDetailRepositoryInPort
	transaksiDetailRepositoryOut TransaksiDetailRepositoryOutPort
}

func NewTransaksiDetailUseCase(TransaksiDetailRepositoryIn TransaksiDetailRepositoryInPort, TransaksiDetailRepositoryOut TransaksiDetailRepositoryOutPort) *TransaksiDetailUseCaseInteractor {
	return &TransaksiDetailUseCaseInteractor{
		transaksiDetailRepositoryIn:  TransaksiDetailRepositoryIn,
		transaksiDetailRepositoryOut: TransaksiDetailRepositoryOut,
	}
}

func (uc *TransaksiDetailUseCaseInteractor) CreateTransaksiDetail(transaksi *domain.TransaksiDetail) (*domain.TransaksiDetail, error) {
	println("di usecase", transaksi)
	return uc.transaksiDetailRepositoryOut.CreateTransaksiDetailResponse(uc.transaksiDetailRepositoryIn.CreateTransaksiDetail(transaksi))
}

func (uc *TransaksiDetailUseCaseInteractor) CreateTransaksiDetailResponse(transaksi *domain.TransaksiDetail) (*domain.TransaksiDetail, error) {
	println("di usecase details", transaksi)
	return uc.transaksiDetailRepositoryOut.CreateTransaksiDetailResponse(uc.transaksiDetailRepositoryIn.CreateTransaksiDetail(transaksi))
}
