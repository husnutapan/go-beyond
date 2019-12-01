package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	var person2 person
	person2.firstName = "test"
	person2.lastName = "test"
	fmt.Println(person2)

	husnu := person{
		firstName: "husnu",
		lastName:  "tapan",
		age:       25,
		contact: contactInfo{
			email:   "test@gmail.com",
			zipCode: 134,
		},
	}
	husnuPointer := &husnu
	husnuPointer.updateName("husnu1")
	husnu.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
