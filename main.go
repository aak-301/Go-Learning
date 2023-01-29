package main

import (
	"log"

	"github.com/aak-301/myniceprogram/helper"
)

// It is used fo concurrency i Go
// It is used to send data from one package to another

const numPool = 1000

func CalculateValue(intChan chan int) {
	randomNumber := helper.RandomNumber(numPool)
	// writing in intChan(channel)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close((intChan))

	go CalculateValue(intChan)

	// reading from intChan(channel)
	num := <-intChan
	log.Println(num)
}
