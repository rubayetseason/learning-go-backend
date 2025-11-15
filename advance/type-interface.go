package main

import "fmt"

type UserInterface interface {
	PrintDetails()
	RecieveMoney(amount float64) float64
}

type NewUser struct {
	Name  string
	Age   int
	Money float64
}

func (usr NewUser) PrintDetails() {
	fmt.Println("User name ~ ", usr.Name)
	fmt.Println("User age ~ ", usr.Age)
	fmt.Println("User money ~ ", usr.Money)
}

func (user NewUser) RecieveMoney(amount float64) float64 {
	return amount + user.Money
}

func main() {
	var user1 UserInterface
	user1 = NewUser{Name: "Alice", Age: 30, Money: 100.0}

	user1.PrintDetails()
}
