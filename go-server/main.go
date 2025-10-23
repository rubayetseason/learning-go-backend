package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello! Server is running!")
}

type Product struct {
	ID          int     `json:"id"` // renames the json key
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageUrl    string  `json:"imageUrl"`

	// camelCase makes these variables private, other packages cannot access --> postman too
}

var productList []Product // slice of products

func productsHandler(w http.ResponseWriter, r *http.Request) {
	sendResponse(w, productList, http.StatusOK)
}

func createProductHandler(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		http.Error(w, "Please enter a valid json", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)
	sendResponse(w, newProduct, http.StatusCreated)

}

func sendResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func main() {
	mux := http.NewServeMux() // router

	mux.Handle("GET /hello", http.HandlerFunc(helloHandler))

	mux.Handle("GET /products", http.HandlerFunc(productsHandler))

	mux.Handle("GET /create-product", http.HandlerFunc(createProductHandler))

	globalRouter := globalHandler(mux)

	fmt.Println("🚀 Golang server is running on: 5000")

	err := http.ListenAndServe(":5000", globalRouter) // if no error then will get nill else will get the error text

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}

func init() {
	prod1 := Product{
		ID:          1,
		Title:       "Product 1",
		Description: "This is a product 1",
		Price:       10.00,
		ImageUrl:    "https://via.placeholder.com/150",
	}

	prod2 := Product{
		ID:          2,
		Title:       "Product 2",
		Description: "This is a product 2",
		Price:       20.00,
		ImageUrl:    "https://via.placeholder.com/150",
	}

	prod3 := Product{
		ID:          3,
		Title:       "Product 3",
		Description: "This is a product 3",
		Price:       30.00,
		ImageUrl:    "https://via.placeholder.com/150",
	}

	productList = append(productList, prod1, prod2, prod3)
}

func globalHandler(mux *http.ServeMux) http.Handler {
	handleAllRequests := func(w http.ResponseWriter, r *http.Request) {
		// handle CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		// handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		mux.ServeHTTP(w, r)

	}

	return http.HandlerFunc(handleAllRequests)
}
