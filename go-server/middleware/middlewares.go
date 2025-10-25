package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	mngr := Manager{globalMiddlewares: make([]Middleware, 0)} // creates an empty slice of middlewares
	return &mngr                                              //this returns the memory address of the mngr variable (a pointer) so that the same instance can be used outside this function.
}

func (mngr *Manager) Use(middlewares ...Middleware) { // pass one or more middleware functions
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...) // append adds these middlewares to the existing globalMiddlewares list
}


// with method wraps your main handler function with all the middlewares
func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	n := next

	for _, middleware := range middlewares {
		n = middleware(n)
	}

	for _, globalMiddleware := range mngr.globalMiddlewares {
		n = globalMiddleware(n)
	}

	return n
}
