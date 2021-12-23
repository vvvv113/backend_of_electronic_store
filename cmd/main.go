package main

import (
	"backend/internal/infrastructure/mysql"
	"backend/internal/interfaces/handlers"
	"backend/internal/interfaces/repository"
	"backend/internal/usecases/storage/product"
	"backend/logger"
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var config = &mysql.Config{
		IP:       "",
		Port:     "",
		User:     "",
		Password: "",
		Database: "",
	}

	logger.Info.Print("Try to connect to db")
	client, err := mysql.New(config)
	if err != nil {
		logger.Critical.Fatalf("Error during MySQL initialization")
	}

	databaseRepository := repository.New(client)
	applicationStorage := product.New(databaseRepository)

	router := mux.NewRouter()
	handlers.Make(router, applicationStorage)
	srv := &http.Server{
		Addr:    ":30003",
		Handler: router,
	}

	go func() {
		listener := make(chan os.Signal, 1)
		signal.Notify(listener, os.Interrupt, syscall.SIGTERM)
		logger.Info.Println("Received a shutdown signal:", <-listener)

		if err := srv.Shutdown(context.Background()); err != nil && err != http.ErrServerClosed {
			logger.Error.Println("Failed to gracefully shutdown ", err)
		}
	}()

	logger.Info.Println("[*]  Listening...")
	if err := srv.ListenAndServe(); err != nil {
		logger.Error.Println("Failed to listen and serve ", err)
	}

	logger.Critical.Fatal("Server shutdown")
}
