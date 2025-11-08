package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/middleware"
	"ecommerce/rest/modules/products"
	"ecommerce/rest/modules/users"
)

func Serve() {

	cnf := config.GetConfig()

	middlewares := middleware.NewShareMiddleware(cnf)

	productHandler := products.NewHandler(middlewares)
	userHandler := users.NewHandler()

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Bootstrap()
}
