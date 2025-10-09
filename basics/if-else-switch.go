package main

import "fmt"

func main() {
	// if-else statement
	a := 10
	b := 20
	isCoding := false

	if a < 10 {
		fmt.Println("less than triggered")
	} else if a == 10 {
		fmt.Println("equal triggered")
	} else {
		fmt.Println("greater than triggered")
	}

	if a <= 10 {
		fmt.Println("less than or equal triggered")
	} else {
		fmt.Println("greater than or equal triggered")
	}

	// switch statement
	switch a {
	case 10:
		fmt.Println("a is 10")
	case 20:
		fmt.Println("a is 20")
	default:
		fmt.Println("a is not 10 or 20")
	}

	if a == 10 && b == 20 {
		fmt.Println("a is 10 and b is 20")
	}

	if a == 10 || b == 20 {
		fmt.Println("a is 10 or b is 20")
	}

	if isCoding {
		fmt.Println("Coding")
	}

	if !isCoding {
		fmt.Println("Not Coding")
	}

}
