package main

import "fmt"

const (
	firstConst  = 1
	secondConst = 2
)

func main() {
	vars()
	arrs()
}

var i int

//normal vars int,float,string
func vars() {

	fmt.Println(i)
	i = 44
	fmt.Println(i)
	var f float32 = 3.14
	fmt.Println(f)

	firstname := "Eslam Salem"
	fmt.Println("Hello Go :D ", firstname)

	greeting := "Hello"

	greetingPtr := &greeting

	fmt.Println(greetingPtr, *greetingPtr)

	greeting = "Salam :D"

	fmt.Println(greetingPtr, *greetingPtr)

	const pi = 3.1415
	fmt.Println(pi)

	fmt.Println(firstConst, secondConst)
}

//arrays
func arrs() {
	var arr [3]int
	arr[0] = 1
	fmt.Println(arr)
	fmt.Println(arr[1-1])

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	slice := arr2[:]
	slice[2] = 6
	fmt.Println(slice)
	fmt.Println(arr2) //the same because the change is slice applied to arr, slice is a pointer to arr2

	freeslice := []int{3, 6}
	freeslice[0] = 0
	freeslice[1] = 5
	freeslice = append(freeslice, 10, 15) //add new elements
	fmt.Println(freeslice)

	s1 := freeslice[1:]
	s2 := freeslice[:2]
	s3 := freeslice[1:2]
	fmt.Println(s1, s2, s3)

	m := map[string]int{"first": 1, "second": 2}
	fmt.Println(m, m["first"])
	delete(m, "first")
	fmt.Println(m, m["first"])

	//structs
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	u.ID = 1
	u.FirstName = "John"
	u.LastName = "Doe"

	fmt.Println(u.FirstName, u.LastName)

	u2 := user{ID: 1, FirstName: "Eslam"}
	fmt.Println(u2, u2.FirstName)
}
