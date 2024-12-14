package main

import "fmt"

type Speaker interface {
	Speak() string
}

func saySomething(say ...Speaker) {
	for _, s := range say {
		fmt.Println(s.Speak())
	}
}

type cat struct {
}

func (c cat) Speak() string {
	return "Purr Meow"
}

type dog struct {
}

func (d dog) Speak() string {
	return "Woff Woff"
}

type person struct {
	name string
}

func (p person) Speak() string {
	return "Hi, My name is " + p.name + " ."
}

func main() {
	c := cat{}
	d := dog{}
	p := person{name: "Kai"}

	saySomething(c, d, p)
}
