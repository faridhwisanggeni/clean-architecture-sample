package infrastructure

import (
	"context"
	"fmt"

	"github.com/cleanarchitect/pos/domain"
	"github.com/jackc/pgx/v4"
)

type TransactionRepositoryImpl struct {
	dbConn *pgx.Conn
}

func NewTransactionRepository(dbConn *pgx.Conn) *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{
		dbConn: dbConn,
	}
}

func (r *TransactionRepositoryImpl) CreateTransaction(transaction *domain.Transaction) (*domain.Transaction, error) {
	sql := "INSERT INTO transaction(total_amount) VALUES($1) RETURNING id, trx_id, total_amount"

	row := r.dbConn.QueryRow(context.Background(), sql, transaction.Total_amount)

	var createdProduct domain.Transaction
	if err := row.Scan(&createdProduct.ID, &createdProduct.Trx_id, &createdProduct.Total_amount); err != nil {
		return nil, fmt.Errorf("Error creating product: %w", err)
	}

	return &createdProduct, nil
}

func (r *TransactionRepositoryImpl) CreateTransactionDetail(transaction *domain.Transaction, item *domain.Item) (*domain.Transaction, error) {
	sql := "INSERT into transaction_detail(trx_id, quantity) VALUES($1,$2) returning trx_id, quantity "

	row := r.dbConn.QueryRow(context.Background(), sql, transaction.Trx_id, item.Quantity)

	var createdProduct domain.Transaction
	if err := row.Scan(&createdProduct.ID, &createdProduct.Trx_id, &createdProduct.Total_amount); err != nil {
		return nil, fmt.Errorf("Error creating product: %w", err)
	}

	return &createdProduct, nil
}

func (r *TransactionRepositoryImpl) CreateTransactionResponse(data *domain.Transaction, err error) (*domain.Transaction, error) {
	return data, err
}

func (r *TransactionRepositoryImpl) CreateTransactionDetailResponse(data *domain.Transaction, err error) (*domain.Transaction, error) {
	return data, err
}
