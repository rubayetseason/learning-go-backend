package routes

import (
	"net/http"

	"ecommerce/rest/middleware"
	"ecommerce/rest/modules/products"
	"ecommerce/rest/modules/users"
)

func InitRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("POST /create-product", manager.With(http.HandlerFunc(products.CreateProductHandler), middleware.AuthenticateJwt))
	mux.Handle("GET /products", manager.With(http.HandlerFunc(products.GetProductsHandler), middleware.EndPrint))
	mux.Handle("GET /products/{productId}", manager.With(http.HandlerFunc(products.GetProductByIdHandler), middleware.AuthenticateJwt))
	mux.Handle("PUT /products/{productId}", manager.With(http.HandlerFunc(products.UpdateProductHandler), middleware.AuthenticateJwt))
	mux.Handle("DELETE /products/{productId}", manager.With(http.HandlerFunc(products.DeleteProductHandler), middleware.AuthenticateJwt))

	mux.Handle("POST /create-user", manager.With(http.HandlerFunc(users.CreateUserHandler)))
	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(users.LoginUserHandler)))
}
