package users

import (
	"encoding/json"
	"net/http"

	"ecommerce/database"
	"ecommerce/util"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser database.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)

	if err != nil {
		http.Error(w, "Please enter a valid json", http.StatusBadRequest)
		return
	}

	createdUser := database.StoreUser(newUser)

	util.SendResponse(w, createdUser, http.StatusCreated)
}
