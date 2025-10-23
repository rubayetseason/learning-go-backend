package main

import (
	"encoding/json"
	"net/http"
)

func createProductHandler(w http.ResponseWriter, r *http.Request) {
	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		http.Error(w, "Please enter a valid json", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(productList) + 1
	productList = append(productList, newProduct)
	sendResponse(w, newProduct, http.StatusCreated)

}
