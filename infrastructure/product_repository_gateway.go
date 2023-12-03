package infrastructure

import (
	"context"
	"fmt"

	"github.com/cleanarchitect/pos/domain"
	"github.com/jackc/pgx/v4"
)

type ProductRepositoryImpl struct {
	dbConn *pgx.Conn
}

func NewProductRepository(dbConn *pgx.Conn) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		dbConn: dbConn,
	}
}

func (r *ProductRepositoryImpl) GetProducts() ([]*domain.Product, error) {
	sql := "SELECT id, name, price FROM product"

	rows, err := r.dbConn.Query(context.Background(), sql)
	if err != nil {
		return nil, fmt.Errorf("error fetching product: %w", err)
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		var p domain.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, fmt.Errorf("error scanning product: %w", err)
		}
		products = append(products, &p)
	}

	return products, nil
}

func (r *ProductRepositoryImpl) GetProductByID(id string) (*domain.Product, error) {
	sql := "SELECT id, name, price FROM product WHERE id = $1"

	row := r.dbConn.QueryRow(context.Background(), sql, id)

	var p domain.Product
	if err := row.Scan(&p.ID, &p.Name, &p.Price); err != nil {
		return nil, fmt.Errorf("error scanning product: %w", err)
	}

	return &p, nil
}

func (r *ProductRepositoryImpl) CreateProduct(product *domain.Product) (*domain.Product, error) {
	sql := "INSERT INTO product(name, price) VALUES($1, $2) RETURNING id, name, price"

	row := r.dbConn.QueryRow(context.Background(), sql, product.Name, product.Price)

	var createdProduct domain.Product
	if err := row.Scan(&createdProduct.ID, &createdProduct.Name, &createdProduct.Price); err != nil {
		return nil, fmt.Errorf("error creating product: %w", err)
	}

	return &createdProduct, nil
}

func (r *ProductRepositoryImpl) UpdateProduct(id string, name string, price float64) (*domain.Product, error) {
	sql := "UPDATE product SET name=$2, price=$3 WHERE id=$1 RETURNING id, name, price"

	row := r.dbConn.QueryRow(context.Background(), sql, id, name, price)

	var updatedProduct domain.Product
	if err := row.Scan(&updatedProduct.ID, &updatedProduct.Name, &updatedProduct.Price); err != nil {
		return nil, fmt.Errorf("error updating product: %w", err)
	}

	return &updatedProduct, nil
}

func (r *ProductRepositoryImpl) DeleteProduct(id string) error {
	sql := "DELETE FROM product WHERE id = $1"

	_, err := r.dbConn.Exec(context.Background(), sql, id)
	if err != nil {
		return fmt.Errorf("error deleting product: %w", err)
	}

	return nil
}

func (r *ProductRepositoryImpl) GetProductsResponse(data []*domain.Product, err error) ([]*domain.Product, error) {
	return data, err
}

func (r *ProductRepositoryImpl) GetProductByIDResponse(data *domain.Product, err error) (*domain.Product, error) {
	return data, err
}

func (r *ProductRepositoryImpl) CreateProductResponse(data *domain.Product, err error) (*domain.Product, error) {
	return data, err
}

func (r *ProductRepositoryImpl) UpdateProductResponse(data *domain.Product, err error) (*domain.Product, error) {
	return data, err
}

func (r *ProductRepositoryImpl) DeleteProductResponse(err error) error {
	return err
}
