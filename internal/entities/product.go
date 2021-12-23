package entities

type Product struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	CategoryID  int           `json:"category_id"`
	Quantity    int           `json:"quantity"`
	Status      ProductStatus `json:"status"`
	Price       int           `json:"price"`
	Photo       string        `json:"photo"`
}

type ProductStatus int

const (
	Available ProductStatus = iota + 1
	NotAvailable
)

type Category struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}
