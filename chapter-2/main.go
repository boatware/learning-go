package main

import (
	"chapter-2/helpers"
	"log"
)

const numPool = 10

func CalculateValue(intChan chan int) {
	value := helpers.RandomNumber(numPool)
	intChan <- value
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
