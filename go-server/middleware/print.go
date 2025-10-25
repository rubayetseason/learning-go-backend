package middleware

import (
	"log"
	"net/http"
)

func Print(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI, "API route hit...")
		next.ServeHTTP(w, r)
	})
}
