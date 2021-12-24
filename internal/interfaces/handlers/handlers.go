package handlers

import (
	"backend/internal/entities"
	"backend/internal/usecases/storage/product"
	"backend/internal/usecases/storage/user"
	"backend/logger"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Controllers struct {
	product product.Controller
	user    user.Controller
}

const (
	productID = "product_id"
)

func getAllProducts(app product.Controller) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info.Println("Got new request to get all products")
		errorMessage := "Error getting all products"
		result, err := app.GetProducts()
		if err != nil {
			logger.Error.Println(err.Error())
			http.Error(w, errorMessage, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})
}

func getProduct(app product.Controller) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info.Printf("Got new request to get product with id %s", mux.Vars(r)[productID])

		errorMessage := "Error getting product"
		productID, err := strconv.Atoi(mux.Vars(r)[productID])
		if err != nil {
			logger.Error.Printf("Failed to parse product ID. Got %v", err)
			http.Error(w, errorMessage, http.StatusBadRequest)
			return
		}
		result, err := app.GetProduct(productID)
		if err != nil {
			logger.Error.Printf(err.Error())
			http.Error(w, errorMessage, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})
}

func addProduct(app product.Controller) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info.Println("Got new request to add product")

		errorMessage := "Error adding new product"

		newProduct := new(entities.Product)
		err := json.NewDecoder(r.Body).Decode(&newProduct)
		if err != nil {
			logger.Error.Printf("Failed to decode product. Got %v", err)
			http.Error(w, errorMessage, http.StatusBadRequest)
			return
		}
		err = app.AddProduct(*newProduct)
		if err != nil {
			logger.Error.Println(err.Error())
			http.Error(w, errorMessage, http.StatusInternalServerError)
			return
		}
	})
}

func createUser(app user.Controller) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info.Println("Got new request to create user")

		errorMessage := "Error creating new user"

		newUser := new(entities.User)
		err := json.NewDecoder(r.Body).Decode(&newUser)

		if err != nil {
			logger.Error.Printf("Failed to decode user. Got %v", err)
			http.Error(w, errorMessage, http.StatusBadRequest)
			return
		}

		err = app.CreateUser(*newUser)
		if err != nil {
			logger.Error.Println(err.Error())
			http.Error(w, errorMessage, http.StatusInternalServerError)
			return
		}
	})
}

func login(app user.Controller) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info.Println("Got new request to login to account")

		errorMessage := "Error login to account"

		var credentials user.Credentials

		err := json.NewDecoder(r.Body).Decode(&credentials)

		if err != nil {
			logger.Error.Printf("Failed to decode credentials. Got %v", err)
			http.Error(w, errorMessage, http.StatusBadRequest)
			return
		}

		cookie, err := app.Login(credentials)

		if err != nil {
			logger.Error.Println(err.Error())
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cookie)
	})
}

func getProfile(app user.Controller) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error getting profile"
		cookie, err := r.Cookie("user_id")

		if err != nil {
			logger.Error.Printf("Failed to get userID. Got %v", err)
			http.Error(w, errorMessage, http.StatusInternalServerError)
			return
		}

		userID, err := strconv.Atoi(cookie.Value)
		if err != nil {
			logger.Error.Printf("Failed to parse userID. Got %v", err)
			http.Error(w, errorMessage, http.StatusInternalServerError)
			return
		}

		logger.Info.Printf("Got new request to get profile with id %s", cookie.Value)

		result, err := app.GetProfile(userID)

		if err != nil {
			logger.Error.Printf(err.Error())
			http.Error(w, errorMessage, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})
}

func Make(r *mux.Router, productApp product.Controller, userApp user.Controller) {
	apiURI := "/api"
	serviceRouter := r.PathPrefix(apiURI).Subrouter()
	serviceRouter.Handle("/products", getAllProducts(productApp)).Methods("GET")
	serviceRouter.Handle("/products/{product_id}", getProduct(productApp)).Methods("GET")
	serviceRouter.Handle("/products", addProduct(productApp)).Methods("POST")
	serviceRouter.Handle("/users/create", createUser(userApp)).Methods("POST")
	serviceRouter.Handle("/users/login", login(userApp)).Methods("POST")
	serviceRouter.Handle("/users/profile", getProfile(userApp)).Methods("GET")
}
