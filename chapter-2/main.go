package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		BirthDate: time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	log.Println(user.FirstName)
}
