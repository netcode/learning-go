package main

import "fmt"

type Human interface {
	getFullName() string
}

type Person struct {
	firstName string
	lastName  string
	age       int
}

func (p Person) greet() string {
	return "Hello my name is " + p.firstName
}

func (p *Person) itsBirthday() {
	p.age++
}

func greet(p Person) string {
	return "Hello " + p.firstName
}

func (p Person) getFullName() string {
	return p.firstName + " " + p.lastName
}

func getPersonFullName(h Human) string {
	return h.getFullName()
}

func main() {
	person := Person{firstName: "Eslam", lastName: "Salem", age: 32}
	fmt.Println(greet(person))
	fmt.Println(person.greet())
	person.itsBirthday()
	fmt.Println(person)

	fmt.Println(getPersonFullName(person))
}
