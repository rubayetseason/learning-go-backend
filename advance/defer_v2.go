package main

import "fmt"

func calculation() (result int) {

	fmt.Println("⭕ NAMED RETURN FUNCTION ⭕")

	fmt.Println("FIRST -->", result)

	show := func() {
		result = result + 10
		fmt.Println("DEFER --> SECOND -->", result)
	}

	defer show()

	result = 5
	fmt.Println("THIRD -->", result)

	return result
}

func calc() int {

	fmt.Println("⭕ NORMAL FUNCTION ⭕")

	result := 0
	fmt.Println("FIRST -->", result)

	show := func() {
		result = result + 10
		fmt.Println("DEFER --> SECOND -->", result)
	}

	defer show()

	result = 5
	fmt.Println("THIRD -->", result)

	return result
}

func main() {
	a := calculation()
	fmt.Println("named function defer -->", a)

	b := calc()
	fmt.Println("normal function defer -->", b)
}

/*
OUTPUT

⭕ NAMED RETURN FUNCTION ⭕
FIRST --> 0
THIRD --> 5
DEFER --> SECOND --> 15
named function defer --> 15

⭕ NORMAL FUNCTION ⭕
FIRST --> 0
THIRD --> 5
DEFER --> SECOND --> 15
normal function defer --> 5

*/

/*
name return func executation step
1. execute regular functions and logics
2. if defer comes store it, if it involves any parent/sibling vars, save those reference address
3. execute entire function
4. execute defer functions
5. if value changes then use that return value

regular return func executation step
1. execute regular functions and logics
2. if defer comes store it, if it involves any parent/sibling vars, save those reference address
3. execute entire function
4. save the return value for return
5. execute defer functions
6. if value changes then still return the step 4 value

*/
