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

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstNAme() string {
	return m.FirstName
}

func main() {
	// user := User{
	// 	FirstName:  "Aakash",
	// 	LastName:   "Shivanshu",
	// 	PoneNumber: "+91-1234567890",
	// }

	// log.Println(user.FirstName)

	var myVar myStruct
	myVar.FirstName = "Aman"

	myVar2 := myStruct{
		FirstName: "Ashish",
	}

	// log.Println("myVar= ", myVar.FirstName)
	// log.Println("myVar2= ", myVar2.FirstName)

	log.Println("myVar= ", myVar.printFirstNAme())
	log.Println("myVar2= ", myVar2.printFirstNAme())
}
