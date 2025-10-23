package global_handler

import "net/http"

func GlobalHandler(mux *http.ServeMux) http.Handler {
	handleAllRequests := func(w http.ResponseWriter, r *http.Request) {
		// handle CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		// handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		mux.ServeHTTP(w, r)

	}

	return http.HandlerFunc(handleAllRequests)
}
