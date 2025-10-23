package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux() // router

	mux.Handle("GET /products", http.HandlerFunc(getProductsHandler))

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
