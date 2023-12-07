package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cleanarchitect/pos/handler"
	"github.com/cleanarchitect/pos/infrastructure"
	"github.com/cleanarchitect/pos/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func main() {
	dbConfig := "postgresql://postgres:1qAzxSw2@localhost:5432/golang"
	db, err := pgx.Connect(context.Background(), dbConfig)
	if err != nil {
		fmt.Println("Error connectng to Postgre:", err)
		return
	}
	defer db.Close(context.Background())

	//initialize dependecies
	productRepository := infrastructure.NewProductRepository(db) //Implement your repository
	productRepositoryInPort := usecase.ProductRepositoryInPort(productRepository)
	productRepositoryOutPort := usecase.ProductRepositoryOutPort(productRepository)
	productUseCase := usecase.NewProductUseCase(productRepositoryInPort, productRepositoryOutPort)
	productHandler := handler.NewProductHandler(productUseCase)

	transactionRepository := infrastructure.NewTransactionRepository(db) //Implement your repository
	transactionRepositoryInPort := usecase.TransactionRepositoryInPort(transactionRepository)
	transactionRepositoryOutPort := usecase.TransactionRepositoryOutPort(transactionRepository)
	transactionUseCase := usecase.NewTransactionUseCase(transactionRepositoryInPort, transactionRepositoryOutPort, productRepositoryInPort, productRepositoryOutPort)
	transactionHandler := handler.NewTransactionHandler(transactionUseCase)

	//set up HTTP Server
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/products", productHandler.GetProducts)
		api.GET("/products/:id", productHandler.GetProductByID)
		api.POST("/products", productHandler.CreateProduct)
		api.PUT("/products/:id", productHandler.UpdateProduct)
		api.DELETE("/products/:id", productHandler.DeleteProduct)
		api.POST("/transaction", transactionHandler.CreateTransaction)
	}

	server := &http.Server{
		Addr:    ":8083", //change this to your desired port
		Handler: router,
	}

	//gin.SetMode(gin.ReleaseMode)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	gracefulShutdown(server)
}

func gracefulShutdown(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	fmt.Println("\nSHutting down gracefully...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Error shutting down: %v\n", err)
	}
	os.Exit(0)
}
