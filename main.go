package main

import (
	"log"
)

func main() {
	var isTrue bool
	isTrue = false

	if isTrue {
		log.Println("isTrue=", isTrue)
	} else {
		log.Println("isTrue=", isTrue)
	}

	cat := "cat"

	if cat == "cat" {
		log.Println("Catis cat")
	} else {
		log.Println("Cat is not cat")
	}

	myNum := 50
	isTrue2 := false

	if myNum > 99 && isTrue2 {
		log.Println("myNum>99 && isTure2 == ", isTrue2)
	} else if myNum > 99 && !isTrue2 {
		log.Println("myNum>99 && isTure2 == ", isTrue2)
	} else if myNum < 100 && isTrue2 {
		log.Println("myNum<100 && isTure2 == ", isTrue2)
	} else if myNum < 100 && !isTrue2 {
		log.Println("myNum<100 && isTure2 == ", isTrue2)
	}

	// switch

	myVar := "fisgh"
	switch myVar {
	case "dog":
		log.Println("This is dog")
	case "cat":
		log.Println("This is cat")
	case "fish":
		log.Println("This is fish")
	default:
		log.Println("not found ")
	}

}
