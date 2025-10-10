package main

import "fmt"

func addFromPackage2(num1 int, num2 int) int {
	sum := num1 + num2
	fmt.Println("Hello, from package 2!")
	return sum
}
