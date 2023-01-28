package main

import (
	"log"
)

func main() {
	var myString string
	myString = "Green"

	log.Println("myString=", myString)

	changeUsingPointers(&myString)
	log.Println("after func call myString=", myString)
}

func changeUsingPointers(s *string) {
	log.Println("string s=", s)
	newValue := "Red"
	*s = newValue
}
