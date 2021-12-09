package repository

import (
	"backend/internal/entities"
)

type driver interface {
	Add(product entities.Product) error
	Get(productID int) (entities.Product, error)
	GetAll() ([]entities.Product, error)
}

type database struct {
	d driver
}

func New(dbHandler driver) *database {
	return &database{
		d: dbHandler,
	}
}

func (db *database) InsertProduct(product entities.Product) error {
	product.Status = entities.Available
	return db.d.Add(product)
}

func (db *database) QueryProduct(productID int) (entities.Product, error) {
	return db.d.Get(productID)
}

func (db *database) QueryProducts() ([]entities.Product, error) {
	return db.d.GetAll()
}
