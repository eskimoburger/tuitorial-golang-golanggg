package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       uint16
}

func (p Person) ShowDetail() {
	fmt.Println(p.firstName, p.lastName, "age", p.age)
}
func (p *Person) SetFirstName(name string) {
	p.firstName = name

}

func main() {
	person := Person{
		firstName: "Alice",
		lastName:  "Blob",
		age:       26,
	}
	person.SetFirstName("Fuck")
	person.ShowDetail()
}
