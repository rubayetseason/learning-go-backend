package main

import (
	"encoding/json"
	"net/http"
)

func sendResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}
