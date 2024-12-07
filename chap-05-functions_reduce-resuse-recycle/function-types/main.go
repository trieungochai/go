package main

import "fmt"

// type calc declares calc to be of the func type,
// determining that it takes two integers as arguments and returns a string
type calc func(int, int) string

// func add(i,j int) string has the same signature as the calc type.
// It takes 2 integers as arguments and returns a string stating “Adding i + j = result.
// Functions can be passed to other functions, just like any other type in Go
func add(i, j int) string {
	result := i + j
	return fmt.Sprintf("Added %d + %d = %d", i, j, result)
}

// func calculator(f calc, i, j int) accepts calc as input.
// The calc type, is a function type that has input parameters of int and a return type of string.
// Anything that matches that signature can be passed to the function.
// The func calculator function returns the result of the function of the calc type.
func calculator(f calc, i, j int) {
	fmt.Println(f(i, j))
}

// we call calculator(add, 5, 6).
// We are passing it the add function. add satisfies the signature of the calc func type
func main() {
	calculator(add, 5, 6)
}

// The ability to pass functions as a type is a powerful feature where you can pass functions to other functions if their signatures match the passed-to function’s input parameter.
