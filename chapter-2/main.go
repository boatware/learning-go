package main

import (
	"errors"
	"log"
)

func main() {
	result, err := divide(10.0, 3.0)
	if err != nil {
		panic(err)
	}

	log.Println(result)
}

func divide(a, b float32) (float32, error) {
	var result float32

	if b == 0 {
		return result, errors.New("cannot divide by zero")
	}

	return a / b, nil
}
