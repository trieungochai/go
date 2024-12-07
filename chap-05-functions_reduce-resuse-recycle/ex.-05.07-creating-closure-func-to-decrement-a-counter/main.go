// In this exercise, we are going to create a closure that decrements from a given starting value.
// We will combine what we have learned about passing an argument to an anonymous function and use that knowledge with a closure.
package main

import "fmt"

func decrement(num int) func() int {
	return func() int {
		num--
		return num
	}
}

func main() {
	counter := 4

	decrementor := decrement(counter)

	fmt.Println("Initial counter value:", counter)
	fmt.Println(decrementor())
	fmt.Println(decrementor())
	fmt.Println(decrementor())
	fmt.Println(decrementor())
}
