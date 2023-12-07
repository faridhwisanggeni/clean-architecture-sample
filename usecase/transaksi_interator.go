package usecase

import (
	"github.com/cleanarchitect/post/domain"
)

type TransaksiUseCaseInteractor struct {
	transaksiRepositoryIn  TransaksiRepositoryInPort
	transaksiRepositoryOut TransaksiRepositoryOutPort
}

func NewTransaksiUseCase(TransaksiRepositoryIn TransaksiRepositoryInPort, TransaksiRepositoryOut TransaksiRepositoryOutPort) *TransaksiUseCaseInteractor {

	// var produc ProductUseCaseInteractor
	// prod,err := produc.GetProductByID("1")
	// fm
	return &TransaksiUseCaseInteractor{
		transaksiRepositoryIn:  TransaksiRepositoryIn,
		transaksiRepositoryOut: TransaksiRepositoryOut,
	}
}

func (uc *TransaksiUseCaseInteractor) CreateTransaksi(transaksi *domain.Transaksi) (*domain.Transaksi, error) {
	println("di usecase", transaksi)
	return uc.transaksiRepositoryOut.CreateTransaksiResponse(uc.transaksiRepositoryIn.CreateTransaksi(transaksi))
}

func (uc *TransaksiUseCaseInteractor) CreateTransaksiResponse(transaksi *domain.Transaksi) (*domain.Transaksi, error) {
	println("di usecase details", transaksi)
	return uc.transaksiRepositoryOut.CreateTransaksiResponse(uc.transaksiRepositoryIn.CreateTransaksi(transaksi))
}
