package main

import "fmt"

var a = 10

func main() {
	a := 30

	if a > 15 {
		a := 46
		fmt.Println(a)
	}
	fmt.Println(a)
}
