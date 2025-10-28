package users

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var requestLogin RequestLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestLogin)

	if err != nil {
		http.Error(w, "Invalid request JSON", http.StatusBadRequest)
		return
	}

	// TODO: Check if the user exists and the password is correct
	user := database.CheckUserEmailAndPassword(requestLogin.Email, requestLogin.Password)

	if user == nil {
		util.SendResponse(w, "Invalid email or password", http.StatusOK)
	}

	util.SendResponse(w, user, http.StatusOK)
}
