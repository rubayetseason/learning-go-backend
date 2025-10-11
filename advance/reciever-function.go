package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func (user Person) printUserDetails() {
	fmt.Println("Name: ", user.name)
	fmt.Println("Age: ", user.age)
	fmt.Println("Job: ", user.job)
	fmt.Println("Salary: ", user.salary)
	fmt.Println("this is from reciever function")
}

func (user Person) call(a int) {
	fmt.Println("this is from reciever function as well ~", a)
}

func main() {

	var pers1 Person

	pers1 = Person{
		name:   "Hege",
		age:    45,
		job:    "Teacher",
		salary: 6000,
	}

	pers1.printUserDetails()
	pers1.call(10)

}
