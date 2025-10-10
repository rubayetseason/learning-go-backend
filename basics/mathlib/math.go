package mathlib

import "fmt"

var Money = 100

func AddFromMathLib(num1 int, num2 int) {
	sum := num1 + num2
	fmt.Println("Hello, from math library!")
	fmt.Println("The sum is ~ ", sum)
}
