package main

import "fmt"

// pass by value --> a full copy of the array is made --> changes inside the function won’t affect the original
// pass by reference --> the function receives a reference to the original array rather than a copy. This is useful when you want to avoid copying large arrays or when you need to modify the original data

func print(arr *[3]int) {
	fmt.Println(arr)
}

func main() {
	arr := [3]int{1, 2, 3}
	print(&arr)
}
