package products

import (
	"ecommerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("POST /create-product", manager.With(http.HandlerFunc(h.CreateProductHandler), middleware.AuthenticateJwt))
	mux.Handle("GET /products", manager.With(http.HandlerFunc(h.GetProductsHandler), middleware.EndPrint))
	mux.Handle("GET /products/{productId}", manager.With(http.HandlerFunc(h.GetProductByIdHandler), middleware.AuthenticateJwt))
	mux.Handle("PUT /products/{productId}", manager.With(http.HandlerFunc(h.UpdateProductHandler), middleware.AuthenticateJwt))
	mux.Handle("DELETE /products/{productId}", manager.With(http.HandlerFunc(h.DeleteProductHandler), middleware.AuthenticateJwt))
}
