package main

import "fmt"

var i int
func main() {

	fmt.Println(i)
	i = 44
	fmt.Println(i)
	var f float32 = 3.14
	fmt.Println(f)

	firstname := "Eslam Salem"
	fmt.Println("Hello Go :D ",firstname);

	greeting := "Hello"

	greetingPtr := &greeting
	
	fmt.Println(greetingPtr, *greetingPtr);

	greeting = "Salam :D"

	fmt.Println(greetingPtr, *greetingPtr);
}