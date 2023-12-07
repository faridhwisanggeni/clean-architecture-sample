package usecase

import (
	"fmt"

	"github.com/cleanarchitect/pos/domain"
)

type TransactionUseCaseInteractor struct {
	transactionRepositoryIn  TransactionRepositoryInPort
	transactionRepositoryOut TransactionRepositoryOutPort
	productRepositoryIn      ProductRepositoryInPort
	productRepositoryOut     ProductRepositoryOutPort
}

func NewTransactionUseCase(transactionRepositoryIn TransactionRepositoryInPort, transactionRepositoryOut TransactionRepositoryOutPort, productRepositoryIn ProductRepositoryInPort, productRepositoryOut ProductRepositoryOutPort) *TransactionUseCaseInteractor {
	return &TransactionUseCaseInteractor{
		transactionRepositoryIn:  transactionRepositoryIn,
		transactionRepositoryOut: transactionRepositoryOut,
		productRepositoryIn:      productRepositoryIn,
		productRepositoryOut:     productRepositoryOut,
	}
}

func (uc *TransactionUseCaseInteractor) CreateTransaction(transaction *domain.Transaction) (*domain.Transaction, error) {
	tot_amt := 0
	for _, v := range transaction.Item {
		//get product detail
		product, _ := uc.productRepositoryOut.GetProductByIDResponse(uc.productRepositoryIn.GetProductByID(v.ID))
		if product.Quantity < v.Quantity {
			var transaction = &domain.Transaction{}
			return uc.transactionRepositoryOut.CreateTransactionResponse(transaction, fmt.Errorf("quantity kebanyakan, stok ga cukup"))
		}
		tot_amt += int(product.Price) * v.Quantity
	}

	transaction.Total_amount = float64(tot_amt)
	newtrx, _ := uc.transactionRepositoryOut.CreateTransactionResponse(uc.transactionRepositoryIn.CreateTransaction(transaction))

	for _, v := range transaction.Item {
		var item = &domain.Item{}
		item.Quantity = v.Quantity
		transaction.Trx_id = newtrx.Trx_id

		uc.transactionRepositoryOut.CreateTransactionDetailResponse(uc.transactionRepositoryIn.CreateTransactionDetail(transaction, item))
	}
	return uc.transactionRepositoryOut.CreateTransactionResponse(transaction, nil)
}

func (uc *TransactionUseCaseInteractor) CreateTransactionResponse(transaction *domain.Transaction) (*domain.Transaction, error) {
	return uc.transactionRepositoryOut.CreateTransactionResponse(uc.transactionRepositoryIn.CreateTransaction(transaction))
}

// CreateTransactionDetail implements TransactionRepositoryInPort.
func (*TransactionUseCaseInteractor) CreateTransactionDetail(*domain.Transaction, *domain.Item) (*domain.Transaction, error) {
	panic("unimplemented")
}

func (uc *TransactionUseCaseInteractor) CreateTransactionDetailResponse(*domain.Transaction, *domain.Item) (*domain.Transaction, error) {
	panic("unimplemented")
}
