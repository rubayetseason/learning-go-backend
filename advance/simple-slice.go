package main

import "fmt"

func main() {
	var arr = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr) // 1, 2, 3, 4, 5, 6

	// array to slice
	slice := arr[1:4] // [2, 3, 4], length 3, capacity 5
	fmt.Println("The slice is ~ ", slice, "with length ~ ", len(slice), "and capacity ~ ", cap(slice))

	// slice to slice
	slice2 := slice[1:3] // [3, 4], length 2, capacity 4 --> cause slice1 is a reference of the main array
	fmt.Println("The slice is ~ ", slice2, "with length ~ ", len(slice2), "and capacity ~ ", cap(slice2))

}

/*
slice --> pointer, length (length of slice), capacity (capacity of slice to the last element of array)
*/
