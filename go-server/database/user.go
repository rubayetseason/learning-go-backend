package database

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"isShopOwner"`
}

var users []User

func StoreUser(user User) User {
	if user.ID != 0 {
		return user
	}

	user.ID = len(users) + 1

	users = append(users, user)
	return user
}

func CheckUserEmailAndPassword(email, password string) *User {
	for _, user := range users {
		if user.Email == email {
			return &user
		}
	}
	return nil
}
