package main

import (
	"fmt"

	"example.com/mathlib"
)

func main() {
	var a int = 10
	fmt.Println("custom package usage!")
	mathlib.AddFromMathLib(mathlib.Money, a)
}

// go mod init example.com
