package products

import (
	"ecommerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("POST /create-product", manager.With(http.HandlerFunc(h.CreateProductHandler), h.middlewares.AuthenticateJwt))
	mux.Handle("GET /products", manager.With(http.HandlerFunc(h.GetProductsHandler), h.middlewares.AuthenticateJwt))
	mux.Handle("GET /products/{productId}", manager.With(http.HandlerFunc(h.GetProductByIdHandler), h.middlewares.AuthenticateJwt))
	mux.Handle("PUT /products/{productId}", manager.With(http.HandlerFunc(h.UpdateProductHandler), h.middlewares.AuthenticateJwt))
	mux.Handle("DELETE /products/{productId}", manager.With(http.HandlerFunc(h.DeleteProductHandler), h.middlewares.AuthenticateJwt))
}
