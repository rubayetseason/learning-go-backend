package main

import (
	"fmt"
	"sync"
)

func printHello(a int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello World", a)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go printHello(i, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish
}
