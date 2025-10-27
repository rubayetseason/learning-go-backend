package routes

import (
	"net/http"

	"ecommerce/rest/middleware"
	"ecommerce/rest/modules/products"
)

func InitRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /products", manager.With(http.HandlerFunc(products.GetProductsHandler), middleware.EndPrint))
	mux.Handle("POST /create-product", manager.With(http.HandlerFunc(products.CreateProductHandler)))
	mux.Handle("GET /products/{productId}", manager.With(http.HandlerFunc(products.GetProductByIdHandler)))
}
