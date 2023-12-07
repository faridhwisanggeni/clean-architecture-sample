package usecase

import "github.com/cleanarchitect/pos/domain"

type TransactionDetailRepositoryInPort interface {
	GetTransactionDetail() ([]*domain.TransactionDetail, error)
	GetTransactionDetailByTrxID(id string) ([]*domain.TransactionDetail, error)
	CreateTransactionDetail(trx *domain.TransactionDetail) (*domain.TransactionDetail, error)
	// UpdateTransaction(product *domain.Product) (*domain.Product, error) //TODO
	// DeleteTransaction(id string) error //TODO
}

type TransactionDetailRepositoryOutPort interface {
	GetTransactionDetailResponse([]*domain.TransactionDetail, error) ([]*domain.TransactionDetail, error)
	GetTransactionDetailByTrxIDResponse([]*domain.TransactionDetail, error) ([]*domain.TransactionDetail, error)
	CreateTransactionDetailResponse(*domain.TransactionDetail, error) (*domain.TransactionDetail, error)
	// UpdateTransactionResponse(*domain.Product, error) (*domain.Product, error) //TODO
	// DeleteTransactionResponse(error) error //TODO
}
