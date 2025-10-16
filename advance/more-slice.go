package main

import "fmt"

func main() {
	// s := []int{1, 2, 3, 4, 5}
	// fmt.Println("The slice is ~ ", s, "with length ~ ", len(s), "and capacity ~ ", cap(s))

	// s := make([]int, 3) // make function is used to create slice
	// s[0] = 1
	// s[1] = 2
	// fmt.Println("The slice is ~ ", s, "with length ~ ", len(s), "and capacity ~ ", cap(s))

	s := make([]int, 3, 5) // make function is used to create slice ==> type, length, capacity
	s[0] = 1
	s[1] = 2
	fmt.Println("The slice is ~ ", s, "with length ~ ", len(s), "and capacity ~ ", cap(s))
}
