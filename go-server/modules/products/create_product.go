package products

import (
	"encoding/json"
	"net/http"

	"ecommerce/database"
	"ecommerce/util"
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		http.Error(w, "Please enter a valid json", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, newProduct)
	util.SendResponse(w, newProduct, http.StatusCreated)

}
