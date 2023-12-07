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
	dbConfig := "postgresql://postgres:0oKmnJi9@localhost:5432/postgres"
	db, err := pgx.Connect(context.Background(), dbConfig)
	if err != nil {
		fmt.Println("error connecting db postgre: ", err)
	}

	defer db.Close(context.Background())

	//Initialize dependencies
	productRepository := infrastructure.NewProductRepository(db)
	transactionRepository := infrastructure.NewTransactionRepository(db)
	transactionDetailRepository := infrastructure.NewTransactionDetailRepository(db)
	transactionRepoInPort := usecase.TransactionRepositoryInPort(transactionRepository)
	transactionRepoOutPort := usecase.TransactionRepositoryOutPort(transactionRepository)
	transactionDetailRepoInPort := usecase.TransactionDetailRepositoryInPort(transactionDetailRepository)
	transactionDetailRepoOutPort := usecase.TransactionDetailRepositoryOutPort(transactionDetailRepository)
	productRepositoryInPort := usecase.ProductRepositoryInPort(productRepository)
	productRepositoryOutPort := usecase.ProductRepositoryOutPort(productRepository)
	productUseCase := usecase.NewProductUseCase(productRepositoryInPort, productRepositoryOutPort)
	productHandler := handler.NewProductHandler(productUseCase)
	transactionUseCase := usecase.NewTransactionUseCase(transactionRepoInPort, transactionRepoOutPort)
	transactionHandler := handler.NewTransactionHandler(transactionUseCase)
	transactionDetailUseCase := usecase.NewTransactionDetailUseCase(transactionDetailRepoInPort, transactionDetailRepoOutPort)
	transactionDetailHandler := handler.NewTransactionDetailHandler(transactionDetailUseCase)

	//set up HTTP server
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/products", productHandler.GetProducts)
		api.GET("/products/:id", productHandler.GetProductByID)
		api.POST("/products", productHandler.CreateProduct)
		api.PUT("/products/:id", productHandler.UpdateProduct)
		api.DELETE("/products/:id", productHandler.DeleteProduct)
		api.GET("/order", transactionHandler.GetTransaction)
		api.GET("/order/:id", transactionHandler.GetTransactionByID)
		api.POST("/order", transactionHandler.CreateTransaction)
		api.GET("/order-detail", transactionDetailHandler.GetTransactionDetail)
		api.GET("/order-detail/:id", transactionDetailHandler.GetTransactionDetailByTrxID)
		api.POST("/order-detail", transactionDetailHandler.CreateTransactionDetail)
	}

	server := &http.Server{
		Addr:    ":8003",
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

	fmt.Println("\nShutting down gracefully...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Error shuting down: %v", err)
	}
	os.Exit(0)
}
