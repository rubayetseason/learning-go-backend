package users

import (
	"ecommerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("POST /create-user", manager.With(http.HandlerFunc(h.CreateUserHandler)))
	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(h.LoginUserHandler)))
}
