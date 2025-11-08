package products

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")
	id, err := strconv.Atoi(productId)

	if err != nil {
		http.Error(w, "Please enter a valid product id", http.StatusBadRequest)
		return
	}

	database.DeleteProduct(id)

	util.SendResponse(w, "Product deleted successfully", http.StatusOK)
}
