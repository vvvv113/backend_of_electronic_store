package entities

type Product struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Quantity    string        `json:"quantity"`
	Status      ProductStatus `json:"status"`
	Price       int           `json:"price"`
	Photo       string        `json:"photo"`
}

type ProductStatus int

const (
	Available ProductStatus = iota + 1
	NotAvailable
)
