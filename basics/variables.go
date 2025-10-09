package main

import "fmt"

func main() {
	a := 10 // int type
	fmt.Println(a)

	var b int = 20 // int type
	fmt.Println(b)

	c := 40.40
	fmt.Println(c)

	d := "Hello"
	fmt.Println(d)

	e := true
	e = false
	// to reassign value we dont need to give :=
	// cannot assign different type of value as well
	fmt.Println(e)

	const pi = 3.14
	fmt.Println(pi)

}

/*
data types
- numeric (int, float, complex) --> int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64
- boolean --> true, false
- string --> "string"
*/
