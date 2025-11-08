package products

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")
	id, err := strconv.Atoi(productId)

	if err != nil {
		http.Error(w, "Please enter a valid product id", http.StatusBadRequest)
		return
	}

	var updatedProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&updatedProduct)

	if err != nil {
		http.Error(w, "Please enter a valid json", http.StatusBadRequest)
		return
	}

	updatedProduct.ID = id
	database.UpdateProduct(updatedProduct)

	util.SendResponse(w, updatedProduct, http.StatusOK)
}
