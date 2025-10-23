package main

import "net/http"

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	sendResponse(w, productList, http.StatusOK)
}
