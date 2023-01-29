package main

import (
	"log"
)

// Maps are immutable
type User struct {
	FirstAme string
	LastName string
}

func main() {

	myMap := make(map[string]string)
	myMap["dog"] = "Maggi"
	myMap["cat"] = "Julie"
	myMap["dog"] = "Juno"
	log.Println(myMap["dog"])
	log.Println(myMap["cat"])

	myMap2 := make(map[string]int)
	myMap2["bihar"] = 1
	myMap2["UP"] = 2

	log.Println(myMap2["bihar"])
	log.Println(myMap2["UP"])

	userMap := make(map[string]User)
	userData := User{
		FirstAme: "Aman",
		LastName: "Singh",
	}

	userMap["first"] = userData
	log.Println(userMap["first"].FirstAme, userMap["first"].LastName)

}
