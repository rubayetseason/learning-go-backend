package cmd

import (
	"fmt"
	"net/http"

	"ecommerce/middleware"
)

func Server() {
	manager := middleware.NewManager()
	manager.Use(middleware.PreflightHandler,
		middleware.CorsHandler,
		middleware.Logger,
		middleware.Print)

	mux := http.NewServeMux() // router
	wrappedMux := manager.WrappedMux(mux)

	initRoutes(mux, manager)

	fmt.Println("🚀 Golang server is running on: 5000")

	err := http.ListenAndServe(":5000", wrappedMux) // if no error then will get nill else will get the error text

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
