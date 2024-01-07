package main

import "fmt"

type Person struct {
	name string
	age  int
	Pet
}

type Pet struct {
	petName string
}

func (person Person) greet() {
	fmt.Println("Hi, my name is " + person.name)
}

func main() {

	var person Person = Person{"Alperen", 21, Pet{"Buddy"}}
	person.greet()
	fmt.Println(person.petName)
}
