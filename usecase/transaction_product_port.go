package usecase

import "github.com/cleanarchitect/pos/domain"

type TransactionRepositoryInPort interface {
	GetTransaction() ([]*domain.TransactionProduct, error)
	GetTransactionByID(id string) (*domain.TransactionProduct, error)
	CreateTransaction(trx *domain.TransactionProduct) (*domain.TransactionProduct, error)
	// UpdateTransaction(product *domain.Product) (*domain.Product, error) //TODO
	// DeleteTransaction(id string) error //TODO
}

type TransactionRepositoryOutPort interface {
	GetTransactionResponse([]*domain.TransactionProduct, error) ([]*domain.TransactionProduct, error)
	GetTransactionByIDResponse(*domain.TransactionProduct, error) (*domain.TransactionProduct, error)
	CreateTransactionResponse(*domain.TransactionProduct, error) (*domain.TransactionProduct, error)
	// UpdateTransactionResponse(*domain.Product, error) (*domain.Product, error) //TODO
	// DeleteTransactionResponse(error) error //TODO
}
