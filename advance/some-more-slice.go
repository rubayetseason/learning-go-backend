package main

import "fmt"

func main() {
	var s []int
	s = append(s, 1, 2, 3)

	var x []int
	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)

	y := x

	x = append(x, 4)
	y = append(y, 5)

	x[0] = 10

	fmt.Println(x) // [10, 2, 3, 5]
	fmt.Println(y) // [10, 2, 3, 5]
}
