package infrastructure

import (
	"context"
	"fmt"

	"strconv"

	"github.com/cleanarchitect/pos/domain"
	"github.com/jackc/pgx/v4"
)

type TransactionRepositoryImpl struct {
	dbConn *pgx.Conn
}

func NewTransactionRepository(dbconn *pgx.Conn) *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{
		dbConn: dbconn,
	}
}

func (r *TransactionRepositoryImpl) GetTransaction() ([]*domain.TransactionProduct, error) {
	sql := "SELECT id, transaction_created_date, product_price_total, product_total_qty FROM transaction_product"

	rows, err := r.dbConn.Query(context.Background(), sql)
	if err != nil {
		return nil, fmt.Errorf("error fetching transactions: %w", err)
	}

	defer rows.Close()

	var transactions []*domain.TransactionProduct
	for rows.Next() {
		var trx domain.TransactionProduct
		if err := rows.Scan(&trx.ID, &trx.ProductPriceTotal, &trx.ProductQtyTotal); err != nil {
			return nil, fmt.Errorf("error fetching transactions: %w", err)
		}
		transactions = append(transactions, &trx)
	}

	return transactions, nil
}

func (r *TransactionRepositoryImpl) GetTransactionByID(id string) (*domain.TransactionProduct, error) {
	sql := "SELECT id, product_price_total, product_total_qty FROM transaction_product where id = $1"

	idconv, _ := strconv.Atoi(id)
	row := r.dbConn.QueryRow(context.Background(), sql, idconv)

	var trx domain.TransactionProduct
	if err := row.Scan(&trx.ID, &trx.ProductPriceTotal, &trx.ProductQtyTotal); err != nil {
		return nil, fmt.Errorf("error scanning transaction: %w", err)
	}

	temp, _ := getDetailByTrxId(r, idconv)
	trx.ProductDetail = temp

	return &trx, nil
}

func getDetailByTrxId(r *TransactionRepositoryImpl, id int) ([]domain.TransactionDetail, error) {
	sql := "SELECT product_id, product_price, product_qty FROM transaction_detail where transaction_id = $1"
	rows, _ := r.dbConn.Query(context.Background(), sql, id)
	defer rows.Close()

	var transactions []domain.TransactionDetail
	for rows.Next() {
		var trx domain.TransactionDetail
		if err := rows.Scan(&trx.ProductID, &trx.ProductPrice, &trx.ProductQty); err != nil {
			return nil, fmt.Errorf("error fetching transactions: %w", err)
		}
		transactions = append(transactions, trx)
	}
	return transactions, nil
}

func (r *TransactionRepositoryImpl) CreateTransaction(trx *domain.TransactionProduct) (*domain.TransactionProduct, error) {
	sql := "INSERT INTO transaction_product(product_price_total, product_total_qty) VALUES($1,$2) RETURNING id"

	row := r.dbConn.QueryRow(context.Background(), sql, trx.ProductPriceTotal, trx.ProductQtyTotal)

	var createdTrx domain.TransactionProduct
	if err := row.Scan(&createdTrx.ID); err != nil {
		fmt.Println("error creating product: ", err)
		return nil, fmt.Errorf("error creating product: %w", err)
	}

	for _, v := range trx.ProductDetail {
		v.TransactionID = createdTrx.ID
		ins, _ := insertDetail(r, &v)
		createdTrx.ProductPriceTotal = createdTrx.ProductPriceTotal + ins.ProductPrice
		createdTrx.ProductQtyTotal = createdTrx.ProductQtyTotal + ins.ProductQty
		createdTrx.ProductDetail = append(createdTrx.ProductDetail, v)
	}

	updateTrxAfterDetailInsert(r, &createdTrx)

	return &createdTrx, nil
}

func insertDetail(r *TransactionRepositoryImpl, trx *domain.TransactionDetail) (*domain.TransactionDetail, error) {
	sql := "INSERT INTO transaction_detail(transaction_id, product_id, product_price, product_qty) VALUES($1,$2,$3,$4) RETURNING transaction_id, product_price, product_qty"

	row := r.dbConn.QueryRow(context.Background(), sql, trx.TransactionID, trx.ProductID, trx.ProductPrice, trx.ProductQty)

	var createdTrx domain.TransactionDetail
	if err := row.Scan(&createdTrx.ID, &createdTrx.ProductPrice, &createdTrx.ProductQty); err != nil {
		fmt.Println("error creating product: ", err)
		return nil, fmt.Errorf("error creating product: %w", err)
	}

	return &createdTrx, nil
}

func updateTrxAfterDetailInsert(r *TransactionRepositoryImpl, trx *domain.TransactionProduct) (*domain.TransactionProduct, error) {
	sql := "UPDATE transaction_product SET product_price_total = $1, product_total_qty = $2 WHERE id=$3 RETURNING id"

	row := r.dbConn.QueryRow(context.Background(), sql, trx.ProductPriceTotal, trx.ProductQtyTotal, trx.ID)

	var createdTrx domain.TransactionProduct
	if err := row.Scan(&createdTrx.ID); err != nil {
		fmt.Println("error updating product: ", err)
		return nil, fmt.Errorf("error updating product: %w", err)
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

func (r *TransactionRepositoryImpl) GetTransactionResponse(data []*domain.TransactionProduct, err error) ([]*domain.TransactionProduct, error) {
	return data, err
}

func (r *TransactionRepositoryImpl) GetTransactionByIDResponse(data *domain.TransactionProduct, err error) (*domain.TransactionProduct, error) {
	return data, err
}

func (r *TransactionRepositoryImpl) CreateTransactionResponse(data *domain.TransactionProduct, err error) (*domain.TransactionProduct, error) {
	return data, err
}

//TODO
// func (r *ProductRepositoryImpl) UpdateProductResponse(data *domain.Product, err error) (*domain.Product, error) {
// 	return data, err
// }

// func (r *ProductRepositoryImpl) DeleteProductResponse(err error) error {
// 	return err
// }
