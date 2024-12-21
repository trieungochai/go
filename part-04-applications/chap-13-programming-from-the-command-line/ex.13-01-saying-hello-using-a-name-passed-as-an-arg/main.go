package main

import (
	"fmt"
	"os"
)

func main() {
	// utilize the args slice to capture the command-line arguments
	arg := os.Args

	// perform validation on the number of arguments supplied,
	// excluding the executable name provided
	if len(arg) < 2 {
		fmt.Println("Usage: go run main.go <firstName> <lastName>")
	}

	// extract the name from the arguments
	firstName := arg[1]
	lastName := arg[2]

	// display a personalized greeting msg
	greeting := fmt.Sprintf("Hello, %s %s! Welcome to the command-line.", firstName, lastName)

	fmt.Println(greeting)
}
