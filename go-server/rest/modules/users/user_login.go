package users

import (
	"ecommerce/config"
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
		return
	}

	cnf := config.GetConfig()

	accessToken, err := util.CreateJwt(cnf.JwtSecret, util.Payload{
		Sub:         user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		IsShopOwner: user.IsShopOwner,
	})

	if err != nil {
		http.Error(w, "Failed to create access token", http.StatusInternalServerError)
		return
	}

	util.SendResponse(w, accessToken, http.StatusOK)
}
