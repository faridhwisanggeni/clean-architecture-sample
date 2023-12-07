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
	sql := "SELECT id, name, price, quantity from product"

	rows, err := r.dbConn.Query(context.Background(), sql)
	if err != nil {
		return nil, fmt.Errorf("Error fetching product: %w", err)
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		var p domain.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Quantity); err != nil {
			return nil, fmt.Errorf("Error scanning product: %w", err)
		}
		products = append(products, &p)
	}

	return products, nil
}

func (r *ProductRepositoryImpl) GetProductByID(id string) (*domain.Product, error) {
	sql := "SELECT id, name, price, quantity from product where id = $1"

	row := r.dbConn.QueryRow(context.Background(), sql, id)

	var p domain.Product
	if err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Quantity); err != nil {
		return nil, fmt.Errorf("Error scanning product: %w", err)
	}

	return &p, nil
}

func (r *ProductRepositoryImpl) CreateProduct(product *domain.Product) (*domain.Product, error) {
	sql := "INSERT INTO product(name, price, quantity) VALUES($1, $2, $3) RETURNING id, name, price, quantity"

	row := r.dbConn.QueryRow(context.Background(), sql, product.Name, product.Price, product.Quantity)

	var createdProduct domain.Product
	if err := row.Scan(&createdProduct.ID, &createdProduct.Name, &createdProduct.Price, &createdProduct.Quantity); err != nil {
		return nil, fmt.Errorf("Error creating product: %w", err)
	}

	return &createdProduct, nil
}

func (r *ProductRepositoryImpl) UpdateProduct(id string, name string, price float64, quantity int) (*domain.Product, error) {
	//cek exist
	var enough int
	sqlcek := "select count(*) from product where id = $1"
	r.dbConn.QueryRow(context.Background(), sqlcek, id).Scan(enough)
	if enough < 1 {
		return nil, fmt.Errorf("Error updating product: %s", "product not found")
	}

	sql := "UPDATE product set name=$2, price=$3, quantity=$4 where id=$1 returning id, name, price, quantity"

	row := r.dbConn.QueryRow(context.Background(), sql, id, name, price, quantity)

	var updatedProduct domain.Product
	if err := row.Scan(&updatedProduct.ID, &updatedProduct.Name, &updatedProduct.Price, &updatedProduct.Quantity); err != nil {
		return nil, fmt.Errorf("Error updating product: %w", err)
	}
	return &updatedProduct, nil
}

func (r *ProductRepositoryImpl) DeleteProduct(id string) error {
	//cek exist
	var enough int
	sqlcek := "select count(*) from product where id = $1"
	r.dbConn.QueryRow(context.Background(), sqlcek, id).Scan(enough)
	if enough < 1 {
		return fmt.Errorf("Error deleting product: %s", "product not found")
	}

	sql := "DELETE FROM product id = $1"

	_, err := r.dbConn.Exec(context.Background(), sql, id)
	if err != nil {
		return fmt.Errorf("Error deleting product: %w", err)
	}
	return nil
}

func (r *ProductRepositoryImpl) GetProductResponse(data []*domain.Product, err error) ([]*domain.Product, error) {
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
