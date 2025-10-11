package main

import "fmt"

const a = 100

var b = 10

func outer() func() {

	money1 := 100

	fmt.Println("==== BANK ====")

	show := func() {
		money1 = money1 + a + b
		fmt.Println(money1)
	}

	return show
}

func call() {
	incr1 := outer()
	incr1()
	incr1()

	incr2 := outer()
	incr2()
	incr2()
}

func main() {
	call()
}

func init() {
	fmt.Println("Hello, from init function!")
}
