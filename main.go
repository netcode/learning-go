package main

import (
	"fmt"

	"github.com/netcode/learning-go/models"
)

func main() {
	u := models.Users{
		ID:        2,
		FirstName: "x",
	}
	fmt.Println("Hello", u)
}
