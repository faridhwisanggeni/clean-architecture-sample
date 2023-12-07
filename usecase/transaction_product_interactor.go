package usecase

import (
	"github.com/cleanarchitect/pos/domain"
)

type TransactionProductUseCaseInteractor struct {
	transactionRepositoryIn  TransactionRepositoryInPort
	transactionRepositoryOut TransactionRepositoryOutPort
}

func NewTransactionUseCase(transactionRepositoryIn TransactionRepositoryInPort, transactionRepositoryOut TransactionRepositoryOutPort) *TransactionProductUseCaseInteractor {
	return &TransactionProductUseCaseInteractor{
		transactionRepositoryIn:  transactionRepositoryIn,
		transactionRepositoryOut: transactionRepositoryOut,
	}
}

func (uc *TransactionProductUseCaseInteractor) GetTransaction() ([]*domain.TransactionProduct, error) {
	return uc.transactionRepositoryOut.GetTransactionResponse(uc.transactionRepositoryIn.GetTransaction())
}

func (uc *TransactionProductUseCaseInteractor) GetTransactionByID(id string) (*domain.TransactionProduct, error) {
	return uc.transactionRepositoryOut.GetTransactionByIDResponse(uc.transactionRepositoryIn.GetTransactionByID(id))
}

func (uc *TransactionProductUseCaseInteractor) CreateTransaction(trx *domain.TransactionProduct) (*domain.TransactionProduct, error) {
	return uc.transactionRepositoryOut.CreateTransactionResponse(uc.transactionRepositoryIn.CreateTransaction(trx))
}

// func (uc *ProductUseCaseInteractor) UpdateProduct(product *domain.Product) (*domain.Product, error) {
// 	return uc.productRepositoryOut.UpdateProductResponse(uc.productRepositoryIn.UpdateProduct(product))
// }
// func (uc *ProductUseCaseInteractor) DeleteProduct(id string) error {
// 	return uc.productRepositoryOut.DeleteProductResponse(uc.productRepositoryIn.DeleteProduct(id))
// }
func (uc *TransactionProductUseCaseInteractor) GetTransactionResponse() ([]*domain.TransactionProduct, error) {
	return uc.transactionRepositoryOut.GetTransactionResponse(uc.transactionRepositoryIn.GetTransaction())
}
func (uc *TransactionProductUseCaseInteractor) GetTransactionByIDResponse(id string) (*domain.TransactionProduct, error) {
	return uc.transactionRepositoryOut.GetTransactionByIDResponse(uc.transactionRepositoryIn.GetTransactionByID(id))
}
func (uc *TransactionProductUseCaseInteractor) CreateTransactionResponse(trx *domain.TransactionProduct) (*domain.TransactionProduct, error) {
	return uc.transactionRepositoryOut.CreateTransactionResponse(uc.transactionRepositoryIn.CreateTransaction(trx))
}

// func (uc *ProductUseCaseInteractor) UpdateProductResponse(product *domain.Product) (*domain.Product, error) {
// 	return uc.productRepositoryOut.UpdateProductResponse(uc.productRepositoryIn.UpdateProduct(product))
// }
// func (uc *ProductUseCaseInteractor) DeleteProductResponse(id string) error {
// 	return uc.productRepositoryOut.DeleteProductResponse(uc.productRepositoryIn.DeleteProduct(id))
// }
