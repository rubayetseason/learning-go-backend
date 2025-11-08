package products

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProductByIdHandler(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")
	id, err := strconv.Atoi(productId)

	if err != nil {
		http.Error(w, "Please enter a valid product id", http.StatusBadRequest)
		return
	}

	product := database.GetSingleProduct(id)
	if product == nil {
		util.SendErrorReponse(w, "Product not found", http.StatusNotFound)
	}

	util.SendResponse(w, product, http.StatusOK)
}
