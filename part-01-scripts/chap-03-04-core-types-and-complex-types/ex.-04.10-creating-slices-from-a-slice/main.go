// use the slice range notation to create slices with a variety of initial values.
// Commonly, in real-world code, you need to work with only a small part of a slice or an array.
// The range notation is a quick and straightforward way of getting only the data you need

package main

import "fmt"

func message() string {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m := fmt.Sprintln("First   :", s[0], s[0:1], s[:1])
	m += fmt.Sprintln("Last    :", s[len(s)-1], s[len(s)-1:len(s)], s[len(s)-1:])
	m += fmt.Sprintln("First 5 :", s[:5])
	m += fmt.Sprintln("Last 4  :", s[5:])
	m += fmt.Sprintln("Middle 5:", s[2:7])
	return m
}

func main() {
	fmt.Print(message())
}
