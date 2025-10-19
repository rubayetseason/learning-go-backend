package main

import "fmt"

func a() {
	i := 0
	fmt.Println("FIRST", i)
	defer fmt.Println("SECOND --> DEFER", i)
	i++
	fmt.Println("THIRD", i)
	defer fmt.Println("FOURTH --> DEFER", i)

	// FIRST 0
	// THIRD 1
	// FOURTH --> DEFER 1
	// SECOND --> DEFER 0
}

func main() {
	a()
}
