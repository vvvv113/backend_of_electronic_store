package entities

type Order struct {
	ID        int         `json:"id"`
	UserID    int         `json:"user_id"`
	Status    OrderStatus `json:"status"`
	Items     []Item
	CreatedAt string `json:"created_at"`
}

type Item struct {
	ProductID int

	Quantity int `json:"quantity"`
}

type OrderStatus int

const (
	InProgress OrderStatus = iota + 1
	Done
	Canceled
)
