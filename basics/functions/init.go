// init function -- we cannot call computer calls it automatically

package main

import "fmt"

var a = 10

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("this is value of a inside main function ~ ", a)
}

func init() {
	fmt.Println("Hello, from init function!")
	fmt.Println("this is value of a inside init function ~ ", a)
	a = 15
}

// output
// Hello, from init function!
// this is value of a inside init function ~  10
// Hello, World!
// this is value of a inside main function ~  15
