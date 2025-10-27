package middleware

import (
	"log"
	"net/http"
)

func EndPrint(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI, "API call ended...")
		next.ServeHTTP(w, r)
	})
}
