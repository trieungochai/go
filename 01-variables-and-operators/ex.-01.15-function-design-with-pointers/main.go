package main

import "fmt"

// we’ll create two functions:
// one that accepts a number by value, adds 5 to it, and then prints the number to the console;
func add5Value(count int) {
	count += 5
	fmt.Println("add5Value     :", count)
}

// and another function that accepts a number as a pointer, adds 5 to it, and then prints the number out.
func add5Point(count *int) {
	*count += 5
	fmt.Println("add5Point     :", *count)
}

func main() {
	var count int
	add5Value(count)
	fmt.Println("add5Value post:", count)

	add5Point(&count)
	fmt.Println("add5Point post:", count)
}
