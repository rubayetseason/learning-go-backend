package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, I am Season! I am a web developer. I am learning Golang.")
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "This method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	encoder := json.NewEncoder(w)

	encoder.Encode(productList)

}

func createProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "This method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		http.Error(w, "Please enter a valid json", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)

	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)
	encoder.Encode(newProduct)

}

func main() {
	mux := http.NewServeMux() // router

	mux.HandleFunc("/hello", helloHandler) // hello route

	mux.HandleFunc("/about-me", aboutHandler) // about-me route

	mux.HandleFunc("/products", productsHandler)

	mux.HandleFunc("/create-product", createProductHandler)

	fmt.Println("🚀 Golang server is running on: 5000")

	err := http.ListenAndServe(":5000", mux) // if no error then will get nill else will get the error text

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
