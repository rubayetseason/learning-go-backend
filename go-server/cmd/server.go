package cmd

import (
	"fmt"
	"net/http"

	"ecommerce/global_handler"
	"ecommerce/middleware"
)

func Server() {
	mux := http.NewServeMux() // router

	manager := middleware.NewManager()

	manager.Use(middleware.Logger, middleware.Print)

	initRoutes(mux, manager)

	globalRouter := global_handler.GlobalHandler(mux)

	fmt.Println("🚀 Golang server is running on: 5000")

	err := http.ListenAndServe(":5000", globalRouter) // if no error then will get nill else will get the error text

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
