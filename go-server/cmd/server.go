package cmd

import (
	"fmt"
	"net/http"

	"ecommerce/global_handler"
	"ecommerce/middleware"
	"ecommerce/modules/products"
)

func Server() {
	mux := http.NewServeMux() // router

	manager := middleware.NewManager()

	mux.Handle("GET /products", manager.With(http.HandlerFunc(products.GetProductsHandler), middleware.Logger))
	mux.Handle("POST /create-product", manager.With(http.HandlerFunc(products.CreateProductHandler), middleware.Logger))
	mux.Handle("GET /products/{productId}", manager.With(http.HandlerFunc(products.GetProductByIdHandler), middleware.Logger))

	globalRouter := global_handler.GlobalHandler(mux)

	fmt.Println("🚀 Golang server is running on: 5000")

	err := http.ListenAndServe(":5000", globalRouter) // if no error then will get nill else will get the error text

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
