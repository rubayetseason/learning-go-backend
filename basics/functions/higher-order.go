package main

import "fmt"

// higher order function that takes function as params
func processOperation(a int, b int, op func(p int, q int)) {
	op(a, b) // callback function
}

// higher order function that returns function
func call() func(x int, y int) {
	return add
}

func add(x int, y int) {
	fmt.Println("The sum is ~ ", x+y)
}

func main() {
	processOperation(10, 20, add)

	sum := call()
	sum(10, 20)
}

// higher order function is a function which accepts a function as an argument or returns a function
