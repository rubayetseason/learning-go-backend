package main

import "fmt"

func main() {

	var a int = 10
	var b int = 20

	fmt.Println("Hello, from package 1!")
	var c = addFromPackage2(a, b)
	fmt.Println(c)
}
