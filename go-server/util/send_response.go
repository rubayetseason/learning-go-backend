package util

import (
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}


func SendErrorReponse(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(message)
}