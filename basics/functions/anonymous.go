package main

import "fmt"

func main() {
	// anonymous function
	//immediately invoked function expression --? IIFE
	func(a, b int) {
		fmt.Println(a + b)
	}(10, 20)
}
