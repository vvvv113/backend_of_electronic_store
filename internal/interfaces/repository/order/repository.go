package order

import (
	"backend/internal/entities"
	storage "backend/internal/usecases/storage/order"
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

func (db *database) InsertOrder(order entities.Order) (int, error) {
	order.Status = entities.InProgress
	return order.ID, db.d.Create(&order)
}

func (db *database) QueryOrder(orderID int) (storage.OrderWithItems, error) {
	var order entities.Order
	err := db.d.FindByID(orderID, order)
	if err != nil {
		logger.Error.Printf("Failed to find order. Error: %s", err)
		return storage.OrderWithItems{}, err
	}

	type Query struct {
		OrderID int
	}

	query := Query{
		OrderID: orderID,
	}

	var items []entities.Item

	err = db.d.FindByParameters(&query, &items, true)
	if err != nil {
		logger.Error.Printf("Failed to find items. Error: %s", err)
		return storage.OrderWithItems{}, err
	}

	return storage.OrderWithItems{Order: order, Items: items}, nil
}

func (db *database) InsertItem(item entities.Item) error {
	return db.d.Create(&item)
}

func (db *database) QueryOrders(userID int) ([]entities.Order, error) {
	var result []entities.Order
	type Query struct {
		UserID int
	}

	query := Query{
		UserID: userID,
	}

	err := db.d.FindByParameters(&query, &result, true)
	if err != nil {
		logger.Error.Printf("Error during getting product: %s", err.Error())
		return []entities.Order{}, err
	}
	return result, nil
}
