package main

import "fmt"

func enterYourName() string {
	var name string
	fmt.Println("Enter your name ~ ")
	fmt.Scanln(&name)
	return "Hello " + name
}

func main() {
	fmt.Println(enterYourName())
}
