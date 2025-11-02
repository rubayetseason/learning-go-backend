package database

import (
	"fmt"
	"strings"
)

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"isShopOwner"`
}

var users []User

func StoreUser(user User) (User, error) {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.Password = strings.TrimSpace(user.Password)

	for _, u := range users {
		if strings.ToLower(u.Email) == user.Email {
			return User{}, fmt.Errorf("email already exists")
		}
	}
	user.ID = len(users) + 1
	users = append(users, user)
	return users[len(users)-1], nil
}

func CheckUserEmailAndPassword(email, password string) *User {
	email = strings.TrimSpace(strings.ToLower(email))
	password = strings.TrimSpace(password)

	for i := range users {
		if strings.ToLower(users[i].Email) == email && users[i].Password == password {
			return &users[i] // pointer to actual element
		}
	}
	return nil
}
