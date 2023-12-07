package usecase

import "github.com/cleanarchitect/pos/domain"

type TransactionDetailUseCaseInteractor struct {
	transactionDetailRepositoryIn  TransactionDetailRepositoryInPort
	transactionDetailRepositoryOut TransactionDetailRepositoryOutPort
}

func NewTransactionDetailUseCase(transactionDetailRepositoryIn TransactionDetailRepositoryInPort, transactionDetailRepositoryOut TransactionDetailRepositoryOutPort) *TransactionDetailUseCaseInteractor {
	return &TransactionDetailUseCaseInteractor{
		transactionDetailRepositoryIn:  transactionDetailRepositoryIn,
		transactionDetailRepositoryOut: transactionDetailRepositoryOut,
	}
}

func (uc *TransactionDetailUseCaseInteractor) GetTransactionDetail() ([]*domain.TransactionDetail, error) {
	return uc.transactionDetailRepositoryOut.GetTransactionDetailResponse(uc.transactionDetailRepositoryIn.GetTransactionDetail())
}

func (uc *TransactionDetailUseCaseInteractor) GetTransactionDetailByTrxID(id string) ([]*domain.TransactionDetail, error) {
	return uc.transactionDetailRepositoryOut.GetTransactionDetailByTrxIDResponse(uc.transactionDetailRepositoryIn.GetTransactionDetailByTrxID(id))
}

func (uc *TransactionDetailUseCaseInteractor) CreateTransactionDetail(trx *domain.TransactionDetail) (*domain.TransactionDetail, error) {
	return uc.transactionDetailRepositoryOut.CreateTransactionDetailResponse(uc.transactionDetailRepositoryIn.CreateTransactionDetail(trx))
}

// func (uc *ProductUseCaseInteractor) UpdateProduct(product *domain.Product) (*domain.Product, error) {
// 	return uc.productRepositoryOut.UpdateProductResponse(uc.productRepositoryIn.UpdateProduct(product))
// }
// func (uc *ProductUseCaseInteractor) DeleteProduct(id string) error {
// 	return uc.productRepositoryOut.DeleteProductResponse(uc.productRepositoryIn.DeleteProduct(id))
// }
func (uc *TransactionDetailUseCaseInteractor) GetTransactionDetailResponse() ([]*domain.TransactionDetail, error) {
	return uc.transactionDetailRepositoryOut.GetTransactionDetailResponse(uc.transactionDetailRepositoryIn.GetTransactionDetail())
}
func (uc *TransactionDetailUseCaseInteractor) GetTransactionDetailByTrxIDResponse(id string) ([]*domain.TransactionDetail, error) {
	return uc.transactionDetailRepositoryOut.GetTransactionDetailByTrxIDResponse(uc.transactionDetailRepositoryIn.GetTransactionDetailByTrxID(id))
}
func (uc *TransactionDetailUseCaseInteractor) CreateTransactionDetailResponse(trx *domain.TransactionDetail) (*domain.TransactionDetail, error) {
	return uc.transactionDetailRepositoryOut.CreateTransactionDetailResponse(uc.transactionDetailRepositoryIn.CreateTransactionDetail(trx))
}

// func (uc *ProductUseCaseInteractor) UpdateProductResponse(product *domain.Product) (*domain.Product, error) {
// 	return uc.productRepositoryOut.UpdateProductResponse(uc.productRepositoryIn.UpdateProduct(product))
// }
// func (uc *ProductUseCaseInteractor) DeleteProductResponse(id string) error {
// 	return uc.productRepositoryOut.DeleteProductResponse(uc.productRepositoryIn.DeleteProduct(id))
// }
