package main

import (
	"fmt"

	"code/chapter5/listing71/entities"
)

func main() {
	u := entities.User{
		Name:  "Bill",
		Email: "bill@email.com",
	}

	fmt.Printf("User: %v\n", u)
}
