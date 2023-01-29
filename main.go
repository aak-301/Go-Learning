package main

import (
	"log"

	"github.com/aak-301/myniceprogram/helpers"
)

func main() {
	log.Println("Hello")

	var myVar helpers.SomeData
	myVar.TypeName = "Some name"
	log.Println(myVar.TypeName)
}
