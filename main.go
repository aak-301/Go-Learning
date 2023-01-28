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
}

func saySomehing() string {
	return "something"
}
