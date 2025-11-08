package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/modules/products"
	"ecommerce/rest/modules/users"
)

func Serve() {

	cnf := config.GetConfig()

	productHandler := products.NewHandler()
	userHandler := users.NewHandler()

	server := rest.NewServer(productHandler, userHandler)
	server.Bootstrap(cnf)
}
