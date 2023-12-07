package unit_tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/cleanarchitect/pos/domain"
	"github.com/cleanarchitect/pos/infrastructure"
	"github.com/cleanarchitect/pos/usecase"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *pgx.Conn {
	dbConfig := "postgresql://postgres:1qAzxSw2@localhost:5432/golang"

	db, err := pgx.Connect(context.Background(), dbConfig)
	if err != nil {
		panic(fmt.Sprintf("Error connecting to psql: %s", err.Error()))
	}

	return db
}

func truncateDB() {
	db := setupTestDB()
	_, err := db.Exec(context.Background(), "TRUNCATE product")
	if err != nil {
		panic(err)
	}
}

func setupInteractor(dbConn *pgx.Conn) *usecase.ProductUseCaseInteractor {
	productRepository := infrastructure.NewProductRepository(dbConn)
	productRepositoryInPort := usecase.ProductRepositoryInPort(productRepository)
	productRepositoryOutPort := usecase.ProductRepositoryOutPort(productRepository)
	productUseCase := usecase.NewProductUseCase(productRepositoryInPort, productRepositoryOutPort)

	return productUseCase
}

func initializeInteractor() *usecase.ProductUseCaseInteractor {
	db := setupTestDB()

	return setupInteractor(db)
}

func TestCreateProductSuccess(t *testing.T) {
	interactor := initializeInteractor()

	data := &domain.Product{
		Name:     "test-first",
		Price:    10,
		Quantity: 4,
	}
	product, err := interactor.CreateProduct(data)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "test-first", product.Name)
	assert.Equal(t, float64(10), product.Price)
	assert.Equal(t, 4, product.Quantity)
	assert.Nil(t, err)
}

func TestCreateProductFailed(t *testing.T) {
	interactor := initializeInteractor()

	t.Run("failed_greaterThan_prouctPrice", func(t *testing.T) {
		data := &domain.Product{
			Name:     "test-first",
			Price:    300,
			Quantity: 4,
		}

		_, err := interactor.CreateProduct(data)
		assert.Error(t, err, "Nilai price a")
	})

	t.Run("failed_greaterThan_prouctQuantity", func(t *testing.T) {
		data := &domain.Product{
			Name:     "test-first",
			Price:    4,
			Quantity: 300,
		}

		_, err := interactor.CreateProduct(data)
		assert.Error(t, err, "Quantity kebanyakan, mobilnya ga muat")
	})
}

func TestUpdateProductSuccess(t *testing.T) {
	interactor := initializeInteractor()

	id := "4e010534-9448-11ee-a832-0242ac110002"

	product, err := interactor.UpdateProduct(
		id,
		"change-test-first",
		11,
		5,
	)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, id, product.ID)
	assert.Equal(t, "change-test-first", product.Name)
	assert.Equal(t, float64(11), product.Price)
	assert.Equal(t, 5, product.Quantity)
	assert.Nil(t, err)
}

func TestUpdateProductFailed(t *testing.T) {
	interactor := initializeInteractor()

	t.Run("id_not_found_update", func(t *testing.T) {
		product, err := interactor.UpdateProduct("36640fa6-943a-11ee-a40d-0242ac110010", "abc", 1, 2)

		assert.Nil(t, product)
		assert.ErrorContains(t, err, "error updating product")
		assert.ErrorContains(t, err, "no rows in result")
	})

	t.Run("invalid_input_id_update", func(t *testing.T) {
		product, err := interactor.UpdateProduct("36640fa6-943a-11ee-a40d-notfound", "abc", 1, 2)

		assert.Nil(t, product)
		assert.ErrorContains(t, err, "error updating product")
		assert.ErrorContains(t, err, "invalid input syntax for type uuid")
	})
}

func TestGetProductByIdSuccess(t *testing.T) {
	interactor := initializeInteractor()

	id := "4e010534-9448-11ee-a832-0242ac110002"

	data, err := interactor.GetProductByID(id)

	assert.Nil(t, err)
	assert.Equal(t, id, data.ID)
	assert.Equal(t, "change-test-first", data.Name)
	assert.Equal(t, float64(11), data.Price)
	assert.Equal(t, 5, data.Quantity)
}

func TestGetProductByIdFailed(t *testing.T) {
	interactor := initializeInteractor()

	t.Run("id_not_found_by_id", func(t *testing.T) {
		data, err := interactor.GetProductByID("36640fa6-943a-11ee-a40d-0242ac110010")

		assert.Nil(t, data)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error scanning product")
		assert.ErrorContains(t, err, "no rows in result set")
	})

	t.Run("invalid_input_id_by_id", func(t *testing.T) {
		product, err := interactor.GetProductByID("36640fa6-943a-11ee-a40d-notfound")

		assert.Nil(t, product)
		assert.ErrorContains(t, err, "error scanning product")
		assert.ErrorContains(t, err, "invalid input syntax for type uuid")
	})
}

func TestDeleteProductSuccess(t *testing.T) {
	interactor := initializeInteractor()

	id := "4e010534-9448-11ee-a832-0242ac110002"

	err := interactor.DeleteProduct(id)
	assert.Nil(t, err)

	data, err := interactor.GetProductByID(id)

	assert.Nil(t, data)
	assert.ErrorContains(t, err, "error scanning product")
	assert.ErrorContains(t, err, "no rows in result")
}

func TestDeleteProductFailed(t *testing.T) {
	interactor := initializeInteractor()

	t.Run("invalid_input_id_delete", func(t *testing.T) {
		err := interactor.DeleteProduct("36640fa6-943a-11ee-a40d-notfound")

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error deleting product")
		assert.ErrorContains(t, err, "invalid input syntax for type uuid")
	})

	t.Run("id_delete_not_found", func(t *testing.T) {
		data, err := interactor.GetProductByID("9f2cd056-9438-11ee-b62a-0242ac110007")
		assert.Nil(t, data)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error scanning product")
		assert.ErrorContains(t, err, "no rows in result set")
	})
}

func TestGetProductsSuccess(t *testing.T) {
	interactor := initializeInteractor()

	datas, err := interactor.GetProducts()

	assert.Nil(t, err)
	assert.Greater(t, len(datas), 0)
}

func TestGetProductsFailed(t *testing.T) {
	truncateDB()
	interactor := initializeInteractor()

	datas, err := interactor.GetProducts()
	if err != nil {
		panic(err)
	}

	assert.LessOrEqual(t, len(datas), 0)
}
