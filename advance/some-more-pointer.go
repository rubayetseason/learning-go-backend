package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Salary float64
}

func main() {
	obj := User{Name: "John", Age: 30, Salary: 5000.0}

	p := &obj
	fmt.Println(&p)       // this is address
	fmt.Println(*p)       // this is value
	fmt.Println(*&p.Name) // this is value
}
