package main

import "fmt"

type Speaker interface {
	Speak() string
}

// declare a cat type and create a method for the cat type called Speak().
// this fulfills the required method sets for the Speaker{} interface.
type cat struct {
}

func (c cat) Speak() string {
	return "Purr Meow"
}

// create a method called chatter that takes the Speaker{} interface as an argument.
func chatter(s Speaker) {
	fmt.Println(s.Speak())
}

// In the main() function,
// we are able to pass a cat type into the chatter function, which can evaluate to the Speaker{} interface.
// This satisfies the required method sets for the interface.
func main() {
	c := cat{}
	chatter(c)
}
