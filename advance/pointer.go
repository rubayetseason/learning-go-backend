package main

import "fmt"

func main() {

	// taking a normal variable
	var x int = 5748

	// declaration of pointer
	var p *int

	// initialization of pointer
	p = &x

	// displaying the result
	fmt.Println("Value stored in x = ", x)
	fmt.Println("Address of x = ", &x)
	fmt.Println("Value stored in variable p = ", p)

	a := 10

	fmt.Println("value before *:", a)

	q := &a
	*q = 20
	fmt.Println("value after *:", a)

	fmt.Println("address:", q)               // address of variable x
	fmt.Println("value at the address:", *q) // value at the address

}
