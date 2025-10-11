package main

import "fmt"

func main() {
	// anonymous function
	//immediately invoked function expression --? IIFE
	func(a, b int) {
		fmt.Println(a + b)
	}(10, 20)

	a := func(a, b int) {
		fmt.Println(a + b)
	}

	a(10, 20)

	// newAdd(1, 2) // wont work because of declrataion before initialization

	newAdd := func(a int, b int) {
		fmt.Println(a + b)
	}

	newAdd(10, 20)

}
