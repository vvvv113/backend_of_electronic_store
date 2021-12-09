package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string        `json:"name"`
	Description string        `json:"description"`
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
