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
	Update(obj interface{}, key string, value string) error
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

func (db *database) QueryOrder(orderID int, userID int) (storage.OrderWithItems, error) {
	var order entities.Order

	type OrderQuery struct {
		ID     int
		UserID int
	}

	orderQuery := OrderQuery{
		ID:     orderID,
		UserID: userID,
	}

	err := db.d.FindByParameters(&orderQuery, &order, false)
	if err != nil {
		logger.Error.Printf("Failed to find order. Error: %s", err)
		return storage.OrderWithItems{}, err
	}

	type Query struct {
		OrderID int
		UserID  int
	}

	itemQuery := Query{
		OrderID: orderID,
	}

	var items []entities.Item

	err = db.d.FindByParameters(&itemQuery, &items, true)
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

func (db *database) UpdateOrder(orderID int, userID int, key string, value string) error {
	var order entities.Order

	type OrderQuery struct {
		ID     int
		UserID int
	}

	orderQuery := OrderQuery{
		ID:     orderID,
		UserID: userID,
	}

	err := db.d.FindByParameters(&orderQuery, &order, false)
	if err != nil {
		logger.Error.Printf("Failed to find order. Error: %s", err)
		return err
	}

	err = db.d.Update(&order, key, value)
	if err != nil {
		logger.Error.Printf("Error during updating parameter %s . Error: %s", key, err)
		return err
	}
	return nil
}
