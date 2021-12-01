package usecases

import (
	"../entities"
)

type repository interface {
	InsertProduct(product entities.Product) error
	QueryProduct(productID string) (entities.Product, error)
	QueryProducts() ([]entities.Product, error)
}

type Controller interface {
	AddProduct(product entities.Product) error
	GetProduct(productID string) (entities.Product, error)
	GetProducts() ([]entities.Product, error)
}

type application struct {
	repo repository
}

func New(repo repository) *application {
	return &application{
		repo: repo,
	}
}

func (app *application) AddProduct(product entities.Product) error {
	return app.repo.InsertProduct(product)
}

func (app *application) GetProduct(productID string) (entities.Product, error) {
	return app.repo.QueryProduct(productID)
}

func (app *application) GetProducts() ([]entities.Product, error) {
	return app.repo.QueryProducts()
}
