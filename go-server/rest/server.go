package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"ecommerce/config"
	"ecommerce/rest/middleware"
	"ecommerce/rest/modules/products"
	"ecommerce/rest/modules/users"
)

type Server struct {
	productHandler *products.Handler
	userHandler    *users.Handler
}

func NewServer(productHandler *products.Handler,
	userHandler *users.Handler) *Server {
	return &Server{
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (server *Server) Bootstrap(cnf config.Config) {

	manager := middleware.NewManager()
	manager.Use(middleware.PreflightHandler,
		middleware.CorsHandler,
		middleware.Logger,
		middleware.Print)

	mux := http.NewServeMux() // router
	wrappedMux := manager.WrappedMux(mux)

	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	portAddress := ":" + strconv.Itoa(cnf.HttpPort)
	fmt.Println("🚀 Golang server is running on port", portAddress)

	err := http.ListenAndServe(portAddress, wrappedMux) // if no error then will get nill else will get the error text

	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}
