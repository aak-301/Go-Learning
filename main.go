package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var whatToSay string
	var i int

	whatToSay = "Enjoy Learning Go"
	fmt.Println(whatToSay)

	i = 10
	fmt.Println(i)

	// := makes the variable dynamic type which detects its datatype on the fly.
	whatWasSaid := saySomehing()
	fmt.Println(whatWasSaid)

	returnMany1, returnMany2 := returnManyThingsFromSameFunction()
	fmt.Println(returnMany1, returnMany2)
}

func saySomehing() string {
	return "something"
}

func returnManyThingsFromSameFunction() (string, string) {
	return "return 1st object\n", "return 2nd object"
}
