package order

import (
	"backend/internal/entities"
	"backend/logger"
)

type repository interface {
	InsertOrder(order entities.Order) (int, error)
	InsertItem(item entities.Item) error
	QueryOrder(orderID int, userID int) (OrderWithItems, error)
	QueryOrders(userID int) ([]entities.Order, error)
}

type Controller interface {
	CreateOrder(userID int, item entities.Item) error
	GetOrders(userID int) ([]entities.Order, error)
	GetOrder(orderID int, userID int) (OrderWithItems, error)
	AddItem(item entities.Item) error
}

type application struct {
	repo repository
}

func New(repo repository) *application {
	return &application{
		repo: repo,
	}
}

type OrderWithItems struct {
	entities.Order
	Items []entities.Item
}

func (app *application) CreateOrder(userID int, item entities.Item) error {
	order := entities.Order{
		UserID: userID,
		Status: entities.InProgress,
	}
	id, err := app.repo.InsertOrder(order)
	if err != nil {
		logger.Error.Printf("Failed to create Order. Error: %s", err)
	}

	item.OrderID = id
	return app.repo.InsertItem(item)
}

func (app *application) GetOrder(orderID int, userID int) (OrderWithItems, error) {
	return app.repo.QueryOrder(orderID, userID)
}

func (app *application) GetOrders(userID int) ([]entities.Order, error) {
	return app.repo.QueryOrders(userID)
}

func (app *application) AddItem(item entities.Item) error {
	return app.repo.InsertItem(item)
}
