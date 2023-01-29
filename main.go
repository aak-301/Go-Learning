package main

import (
	"log"
)

func main() {
	for i := 0; i < 10; i++ {
		log.Println("i =", i)
	}

	animals := []string{"dog", "fish", "cat"}
	for i, animal := range animals {
		log.Println(i, animal)
	}

	for _, animal := range animals {
		log.Println(animal)
	}

	animals2 := make(map[string]string)
	animals2["dog"] = "fido"
	animals2["cat"] = "fluffy"
	for animalType, anianimal := range animals2 {
		log.Println(animalType, anianimal)
	}

	var line = "Hey there. njoy learning GOLang"
	for i, l := range line {
		log.Println(i, l)
	}

	type User struct {
		FirstName string
		LstName   string
		Email     string
		Age       int
	}

	var user []User
	user = append(user, User{"Aman", "Singh", "Aman@mailinaor.com", 20})
	user = append(user, User{"Adarsh", "Singh", "Adarsh@mailinaor.com", 30})
	user = append(user, User{"Shish", "Kumar", "Ashish@mailinaor.com", 10})
	for _, u := range user {
		log.Println(u)
	}
}
