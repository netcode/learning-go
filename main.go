package main

import (
	"fmt"

	"github.com/netcode/learning-go/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "x",
	}
	fmt.Println("Hello", u)
}
