package main

import (
	"log"
	"time"
)

// In Go whe we declare a variable,function,etc with capital letters
// then it can be accessed in different packages as well

type User struct {
	FirstName  string
	LastName   string
	PoneNumber string
	Age        int
	BirthDate  time.Time
}

func main() {
	user := User{
		FirstName:  "Aakash",
		LastName:   "Shivanshu",
		PoneNumber: "+91-1234567890",
	}

	log.Println(user.FirstName)
}
