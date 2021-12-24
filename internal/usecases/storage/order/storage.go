package order

import (
	"backend/internal/entities"
	"backend/logger"
	"fmt"
)

type repository interface {
	InsertOrder(order entities.Order) (int, error)
	InsertItem(item entities.Item) error
	QueryOrder(orderID int, userID int) (OrderWithItems, error)
	QueryOrders(userID int) ([]entities.Order, error)
	UpdateOrder(orderID int, userID int, key string, value string) error
	FindOrderByParam(obj interface{}) (int, error)
}

type Controller interface {
	AddToCart(userID int, item entities.Item) error
	GetOrders(userID int) ([]entities.Order, error)
	GetOrder(orderID int, userID int) (OrderWithItems, error)
	ChangeStatus(orderID int, userID int, status entities.OrderStatus) error
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

type OrderQuery struct {
	UserID int
	Status entities.OrderStatus
}

func (app *application) AddToCart(userID int, item entities.Item) error {
	inProgressOrderID, err := app.repo.FindOrderByParam(OrderQuery{UserID: userID, Status: entities.InProgress})
	if inProgressOrderID == 0 {
		order := entities.Order{
			UserID: userID,
			Status: entities.InProgress,
		}
		item.OrderID, err = app.repo.InsertOrder(order)
		if err != nil {
			logger.Error.Printf("Failed to create Order. Error: %s", err)
		}
	} else {
		item.OrderID = inProgressOrderID
	}

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

func (app *application) ChangeStatus(orderID int, userID int, status entities.OrderStatus) error {
	return app.repo.UpdateOrder(orderID, userID, "status", fmt.Sprintf("%d", status))
}
