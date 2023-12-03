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
	// dsn := "user=g2dev01 password=g2dev01password dbname=g2_database sslmode=disable"
	dbConfig := "postgresql://g2dev01:g2dev01password@localhost:54320/g2_database"
	db, err := pgx.Connect(context.Background(), dbConfig)
	if err != nil {
		fmt.Println("Error connecting to PostgreSQL:", err)
		return
	}
	defer db.Close(context.Background())

	// Initialize dependencies
	productRepository := infrastructure.NewProductRepository(db) // Implement your repository
	productUseCase := usecase.NewProductUseCase(productRepository)
	productHandler := handler.NewProductHandler(productUseCase)

	// Set up HTTP server
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/products", productHandler.GetProducts)
		api.GET("/products/:id", productHandler.GetProductByID)
		api.POST("/products", productHandler.CreateProduct)
		api.PUT("/products/:id", productHandler.UpdateProduct)
		api.DELETE("/products/:id", productHandler.DeleteProduct)
	}

	server := &http.Server{
		Addr:    ":8083", // Change this to your desired port
		Handler: router,
	}

	gin.SetMode(gin.ReleaseMode)

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
		fmt.Printf("Error shutting down: %v\n", err)
	}
	os.Exit(0)
}
