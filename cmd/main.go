package main

import (
	"backend/internal/infrastructure/mysql"
	"encoding/json"
	"fmt"
)

func main() {
	var config = &mysql.Config{
		IP:       "",
		Port:     "",
		User:     "",
		Password: "",
		Database: "",
	}

	client, err := mysql.New(config)
	if err != nil {
		return
	}

	// product := entities.Product{Name: "Iphone XE", Quantity: 10, Description: "test", Status: entities.Available, Price: 10000, Photo: "null"}

	// client.Add(product)
	product, _ := client.Get(1)
	data, err := json.Marshal(product)
	fmt.Println(string(data))

	products, _ := client.GetAll()
	data, err = json.Marshal(products)
	fmt.Println(string(data))
}
