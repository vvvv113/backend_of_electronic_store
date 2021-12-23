package main

import (
	"backend/internal/infrastructure/mysql"
	"backend/internal/interfaces/handlers"
	productRepo "backend/internal/interfaces/repository/product"
	userRepo "backend/internal/interfaces/repository/user"
	"backend/internal/usecases/storage/product"
	"backend/internal/usecases/storage/user"
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
		IP:       "127.0.0.1",
		Port:     "3306",
		User:     "admin",
		Password: "admin",
		Database: "shop",
	}

	logger.Info.Print("Try to connect to db")
	client, err := mysql.New(config)
	if err != nil {
		logger.Critical.Fatalf("Error during MySQL initialization")
	}

	productStorage := product.New(productRepo.New(client))

	userStorage := user.New(userRepo.New(client))

	router := mux.NewRouter()
	handlers.Make(router, productStorage, userStorage)
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
