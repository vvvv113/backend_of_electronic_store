package entities

type Order struct {
	ID     int         `json:"id"`
	UserID int         `json:"user_id"`
	Status OrderStatus `json:"status"`
}

type Item struct {
	OrderID   int `json:"order_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type OrderStatus int

const (
	InProgress OrderStatus = iota + 1
	Done
	Canceled
)
