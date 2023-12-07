package infrastructure

import (
	"context"
	"fmt"
	"github.com/cleanarchitect/pos/domain"

	"strconv"

	"github.com/jackc/pgx/v4"
)

type TransactionDetailRepositoryImpl struct {
	dbConn *pgx.Conn
}

func NewTransactionDetailRepository(dbconn *pgx.Conn) *TransactionDetailRepositoryImpl {
	return &TransactionDetailRepositoryImpl{
		dbConn: dbconn,
	}
}

func (r *TransactionDetailRepositoryImpl) GetTransactionDetail() ([]*domain.TransactionDetail, error) {
	sql := "SELECT transaction_id, product_id, product_price, product_qty FROM transaction_detail"

	rows, err := r.dbConn.Query(context.Background(), sql)
	if err != nil {
		return nil, fmt.Errorf("error fetching transactions: %w", err)
	}

	defer rows.Close()

	var transactionDetails []*domain.TransactionDetail
	for rows.Next() {
		var trx domain.TransactionDetail
		if err := rows.Scan(&trx.TransactionID, &trx.ProductID, &trx.ProductPrice, &trx.ProductQty); err != nil {
			return nil, fmt.Errorf("error fetching transactions: %w", err)
		}
		transactionDetails = append(transactionDetails, &trx)
	}

	return transactionDetails, nil
}

func (r *TransactionDetailRepositoryImpl) GetTransactionDetailByTrxID(trxID string) ([]*domain.TransactionDetail, error) {
	sql := "SELECT transaction_id, product_id, product_price, product_qty FROM transaction_detail where transaction_id = $1"

	idconv, _ := strconv.Atoi(trxID)
	rows, err := r.dbConn.Query(context.Background(), sql, idconv)
	if err != nil {
		return nil, fmt.Errorf("error fetching transactions: %w", err)
	}

	defer rows.Close()

	var transactionDetails []*domain.TransactionDetail
	for rows.Next() {
		var trx domain.TransactionDetail
		if err := rows.Scan(&trx.TransactionID, &trx.ProductID, &trx.ProductPrice, &trx.ProductQty); err != nil {
			return nil, fmt.Errorf("error fetching transactions: %w", err)
		}
		transactionDetails = append(transactionDetails, &trx)
	}

	return transactionDetails, nil
}

func (r *TransactionDetailRepositoryImpl) CreateTransactionDetail(trx *domain.TransactionDetail) (*domain.TransactionDetail, error) {
	sql := "INSERT INTO transaction_detail(transaction_id, product_id, product_price, product_qty) VALUES($1,$2,$3,$4) RETURNING id, product_price, product_qty"

	row := r.dbConn.QueryRow(context.Background(), sql, trx.TransactionID, trx.ProductID, trx.ProductPrice, trx.ProductQty)

	var createdTrx domain.TransactionDetail
	if err := row.Scan(&createdTrx.ID, &createdTrx.ProductPrice, &createdTrx.ProductQty); err != nil {
		return nil, fmt.Errorf("error creating product: %w", err)
	}

	return &createdTrx, nil
}

//TODO
// func (r *ProductRepositoryImpl) UpdateProduct(product *domain.Product) (*domain.Product, error) {
// 	sql := "UPDATE public.product SET name=$2, price=$3, quantity=$4 WHERE id=$1 RETURNING id, name, price, quantity"

// 	row := r.dbConn.QueryRow(context.Background(), sql, product.ID, product.Name, product.Price, product.Quantity)

// 	var updatedProduct domain.Product
// 	if err := row.Scan(&updatedProduct.ID, &updatedProduct.Name, &updatedProduct.Price, &updatedProduct.Quantity); err != nil {
// 		return nil, fmt.Errorf("error updating product: %w", err)
// 	}

// 	return &updatedProduct, nil
// }

// func (r *ProductRepositoryImpl) DeleteProduct(id string) error {
// 	sql := "DELETE FROM public.product WHERE id=$1"

// 	_, err := r.dbConn.Exec(context.Background(), sql, id)

// 	if err != nil {
// 		return fmt.Errorf("error deleting product: %w", err)
// 	}

// 	return nil
// }

func (r *TransactionDetailRepositoryImpl) GetTransactionDetailResponse(data []*domain.TransactionDetail, err error) ([]*domain.TransactionDetail, error) {
	return data, err
}

func (r *TransactionDetailRepositoryImpl) GetTransactionDetailByTrxIDResponse(data []*domain.TransactionDetail, err error) ([]*domain.TransactionDetail, error) {
	return data, err
}

func (r *TransactionDetailRepositoryImpl) CreateTransactionDetailResponse(data *domain.TransactionDetail, err error) (*domain.TransactionDetail, error) {
	return data, err
}

//TODO
// func (r *ProductRepositoryImpl) UpdateProductResponse(data *domain.Product, err error) (*domain.Product, error) {
// 	return data, err
// }

// func (r *ProductRepositoryImpl) DeleteProductResponse(err error) error {
// 	return err
// }
