package products

import (
	"net/http"

	"ecommerce/database"
	"ecommerce/util"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	util.SendResponse(w, database.ListOfProducts(), http.StatusOK)
}
