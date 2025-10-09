package main

import "fmt"

func addWithoutReturn(num1 int, num2 int) {
	sum := num1 + num2
	fmt.Println(sum)
}

func addWithReturn(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func getAdditionAndMultiplication(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	mult := num1 * num2
	return sum, mult
}

func main() {
	a := 10
	b := 20
	addWithoutReturn(a, b)

	c := addWithReturn(a, b)
	fmt.Println(c)

	d, e := getAdditionAndMultiplication(a, b)
	fmt.Println(d, e)
}
