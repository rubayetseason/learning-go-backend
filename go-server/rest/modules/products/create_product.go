package products

import (
	"encoding/json"
	"net/http"

	"ecommerce/database"
	"ecommerce/util"
)

func (h *Handler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		http.Error(w, "Please enter a valid json", http.StatusBadRequest)
		return
	}

	createdProduct := database.StoreProduct(newProduct)

	util.SendResponse(w, createdProduct, http.StatusCreated)

}
