package main

import (
	"log"

	"github.com/aak-301/myniceprogram/helper"
)

const numPool = 1000

func CalculateValue(intChan chan int) {
	randomNumber := helper.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close((intChan))

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
