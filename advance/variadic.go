package main

import "fmt"

func print(numbers ...int) {
	fmt.Println("numbers:", numbers, "length:", len(numbers), "capacity:", cap(numbers))
}

func main() {
	print(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	print(5, 6, 7, 8, 9, 10)
}
