package main

import "fmt"

type Thing interface {
	ShowDetails()
}

type Person struct {
	firstName string
	lastName  string
	age       uint8
}

func (p Person) ShowDetails() {

	fmt.Println(p.firstName, p.lastName, p.age)
}

type Animal struct {
	name string
	age  uint8
}

func (a Animal) ShowDetails() {
	fmt.Println(a.name, a.age)
}

func Main() {
	person := Person{
		firstName: "John",
		lastName:  "Mayer",
		age:       50,
	}

	animal := Animal{
		name: "HeeHa",
		age:  5,
	}

	for _, v := range []Thing{person, animal} {
		v.ShowDetails()
	}
}
