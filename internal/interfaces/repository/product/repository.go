package product

import (
	"backend/internal/entities"
	"backend/logger"
)

type driver interface {
	Create(obj interface{}) error
	FindAll(obj interface{}) error
	FindByParameters(searchObj interface{}, obj interface{}, isAll bool) error
	FindByID(ID int, obj interface{}) error
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
	return db.d.Create(&product)
}

func (db *database) QueryProduct(productID int) (entities.Product, error) {
	var result entities.Product
	err := db.d.FindByID(productID, &result)
	if err != nil {
		logger.Error.Printf("Error during getting product: %s", err.Error())
		return entities.Product{}, err
	}
	return result, nil
}

func (db *database) QueryProducts() ([]entities.Product, error) {
	var result []entities.Product
	err := db.d.FindAll(&result)
	if err != nil {
		logger.Error.Printf("Error during getting products: %s", err.Error())
		return []entities.Product{}, err
	}
	return result, nil
}
