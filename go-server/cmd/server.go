package cmd

import (
	"ecommerce/global_handler"
	"ecommerce/modules/products"
	"fmt"
	"net/http"
)

func Server() {
	mux := http.NewServeMux() // router

	mux.Handle("GET /products", http.HandlerFunc(products.GetProductsHandler))

	mux.Handle("GET /create-product", http.HandlerFunc(products.CreateProductHandler))

	globalRouter := global_handler.GlobalHandler(mux)

	fmt.Println("🚀 Golang server is running on: 5000")

	err := http.ListenAndServe(":5000", globalRouter) // if no error then will get nill else will get the error text

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
