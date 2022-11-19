package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       uint8
}

func main() {
	person := Person{
		firstName: "Alice",
		lastName:  "Bob",
		age:       27}
	fmt.Printf("%v\n", person)
	fmt.Println(person.firstName, person.lastName, "age", person.age)
	person.age += 1
	fmt.Println(person.firstName, person.lastName, "age", person.age)

}
