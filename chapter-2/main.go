// Package decoration.
// "main" is the name of the package.
package main

import "fmt"

// Entry point of the program.
func main() {
	fmt.Println("Hello, World!")

	var whatToSay string
	var i int

	whatToSay = "Hello, World!"
	fmt.Println(whatToSay)

	i = 42
	fmt.Println("The value of i is", i)

	whatWasSaid, whatElseWasSaid := saySomething()
	fmt.Println("The function returned", whatWasSaid, "and", whatElseWasSaid)
}

func saySomething() (string, string) {
	return "Hello, World!", "Bye, World!"
}
