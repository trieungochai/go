// We are going to implement some code that takes a concrete type (cat, dog, or person)
// and satisfies the Speaker{} interface type.

package main

import "fmt"

type Speaker interface {
	Speak() string
}

// We have 3 concrete types (cat, dog, and person).
// The cat and dog types are empty structs, while the person struct has a name field.
type dog struct {
}

type cat struct {
}

type person struct {
	name string
}

// In the main() function, we call catSpeak(), dogSpeak(), and personSpeak() to invoke their respective Speak() methods.
// The code has a lot of redundant functions that perform similar actions.
// We can refactor this code to be simpler and easier to read.
func main() {
	c := cat{}
	d := dog{}
	p := person{name: "Kai"}

	catSpeak(c)
	dogSpeak(d)
	personSpeak(p)
}

// Each of our types implicitly implements the Speaker{} interface.
// Each of the concrete types implements it differently from the others.
func (c cat) Speak() string {
	return "Purr Meow"
}

func (d dog) Speak() string {
	return "Woof Woof"
}

func (p person) Speak() string {
	return "Hi, My name is " + p.name + "."
}

func catSpeak(c cat) {
	fmt.Println(c.Speak())
}

func dogSpeak(d dog) {
	fmt.Println(d.Speak())
}

func personSpeak(p person) {
	fmt.Println(p.Speak())
}
