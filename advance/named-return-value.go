package main

func sum(a int, b int) (s int) { // s is a named return value which returns the var name as return
	s = a + b
	return
}

func main() {
	result := sum(10, 20)
	println(result)
}
