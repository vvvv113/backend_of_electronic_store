package handlers

import (
	"backend/internal/usecases/storage"
	"encoding/json"
	"fmt"
	"net/http"
)

func getAllProducts(app storage.Controller) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error getting all products"
		result, err := app.GetProducts()
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, errorMessage, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})
}
