package main

import (
	"log"
	"sort"
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

	// slices

	var myString string
	myString = "Fish"
	log.Println(myString)

	var myString2 []string
	myString2 = append(myString2, "Aakash")
	myString2 = append(myString2, "Shivanshu")
	log.Println(myString2)

	var myString3 []int
	myString3 = append(myString3, 3)
	myString3 = append(myString3, 1)
	myString3 = append(myString3, 2)
	log.Println(myString3)

	sort.Ints(myString3)
	log.Println(myString3)

	numbers := []int{1, 2, 3, 7, 5, 9, 0}
	log.Println(numbers)
	log.Println(numbers[0:2])
}
