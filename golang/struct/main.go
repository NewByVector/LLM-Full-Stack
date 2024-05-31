package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	var p Person
	p.firstName = "John"
	p.lastName = "Doe"
	p.age = 25

	fmt.Println(p)
}
