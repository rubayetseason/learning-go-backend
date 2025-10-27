package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"ecommerce/config"
	"ecommerce/rest/middleware"
	"ecommerce/rest/routes"
)

func Server(cnf config.Config) {

	manager := middleware.NewManager()
	manager.Use(middleware.PreflightHandler,
		middleware.CorsHandler,
		middleware.Logger,
		middleware.Print)

	mux := http.NewServeMux() // router
	wrappedMux := manager.WrappedMux(mux)

	routes.InitRoutes(mux, manager)

	portAddress := ":" + strconv.Itoa(cnf.HttpPort)
	fmt.Println("🚀 Golang server is running on port", portAddress)

	err := http.ListenAndServe(portAddress, wrappedMux) // if no error then will get nill else will get the error text

	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}
