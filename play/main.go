package main

import (
	"fmt"

	"github.com/netcode/learning-go/play/utils"
)

func greeting(name string) string {
	return "Hi " + name
}

func compareInt(x int32, y int32) {
	if x < y {
		fmt.Printf("%d is less than %d \n", x, y)
	} else if x == y {
		fmt.Printf("%d is equal to %d \n", x, y)
	} else {
		fmt.Printf("%d is bigger than %d \n", x, y)
	}
	return
}

func loops() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Hello", i)
	}

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}

		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}

		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}

		fmt.Println(i)
	}
}

func ranges() {
	ids := []int{33, 44, 55, 66, 77}

	for _, id := range ids {
		fmt.Printf("Hello ID %d\n", id)
	}

	emails := map[string]string{"name": "eslam", "email": "fakeEmail@gmail.com"}

	for k, v := range emails {
		fmt.Println(k, " ", v)
	}
}

func pointers() {
	a := 5
	b := &a
	*b = 6
	fmt.Println(a, *b)
}

func anon(y int) {
	addfun := func(x int) int {
		sum := x + 1
		return sum
	}

	finalNum := addfun(y)
	fmt.Println(finalNum)
}

func main() {
	name := "Eslam"

	fmt.Println(greeting(name))

	x := utils.Reverse("hello")
	y := utils.AddNum(6, 3)
	z := utils.SubNum(6, 3)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	var arr [2]string
	arr[0] = "koko"
	arr[1] = "spsp"
	fmt.Println(arr)

	arr2 := [2]string{"koko", "soso"}
	fmt.Println(arr2)

	arr3 := []int32{1, 2, 3}
	arr3 = append(arr3, 4)
	fmt.Println(arr3)

	compareInt(5, 3)

	loops()

	ranges()

	pointers()

	anon(3)
}
