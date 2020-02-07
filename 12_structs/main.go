package main

import (
	"fmt"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{
		firstName: "Samantha",
		lastName:  "Smith",
		city:      "Boston",
		gender:    "M",
		age:       30,
	}

	fmt.Println(person1.greet())
	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1)
	person1.hasBirthday()
	fmt.Println(person1)
	person1.getMarried("smithy")
	fmt.Println(person1)

}
