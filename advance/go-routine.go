package main

import (
	"fmt"
	"time"
)

var a = 10

const p = 11

func printHello(a int) {
	fmt.Println("Hello World", a)
}

func main() {

	go printHello(1)

	go printHello(2)

	go printHello(3)

	go printHello(4)

	go printHello(5)

	fmt.Println(a, " ", p)

	time.Sleep(5 * time.Second)

}
