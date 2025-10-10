package main

import "fmt"

var (
	a = 10
	b = 20
)

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func main() {
	var c int = 30
	var d int = 40
	fmt.Println(add(a, b))
	fmt.Println(add(c, d))
}
