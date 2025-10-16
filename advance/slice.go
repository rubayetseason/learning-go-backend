package main

import "fmt"

func changeSlice(x []int) []int {
	x[0] = 100
	x = append(x, 200)
	return x
}

func main() {
	x := []int{1, 2, 3, 4, 5}
	x = append(x, 6)
	x = append(x, 7)

	a := x[4:]

	y := changeSlice(a)

	fmt.Println(x)
	fmt.Println(y)
}
