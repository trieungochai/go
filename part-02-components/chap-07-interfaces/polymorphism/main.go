package main

import "fmt"

type Speaker interface {
	Speak() string
}

type cat struct{}

func (c cat) Speak() string {
	return "Purr Meow"
}

func catSpeak(c cat) {
	fmt.Println(c.Speak())
}

func main() {
	c := cat{}
	catSpeak(c)
}

// cat satisfies the Speaker{} interface. The main() function calls catSpeak() and takes a type of cat.
// Inside catSpeak(), it prints out the results of its Speak() method.
