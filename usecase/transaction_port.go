package usecase

import "github.com/cleanarchitect/pos/domain"

type TransactionRepositoryInPort interface {
	CreateTransaction(transaction *domain.Transaction) (*domain.Transaction, error)
	CreateTransactionDetail(*domain.Transaction, *domain.Item) (*domain.Transaction, error)
}

type TransactionRepositoryOutPort interface {
	CreateTransactionResponse(*domain.Transaction, error) (*domain.Transaction, error)
	CreateTransactionDetailResponse(*domain.Transaction, error) (*domain.Transaction, error)
}
