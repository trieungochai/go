// The three dots (…) in front of the type are called a pack operator.
// The pack operator is what makes it a variadic function.
// It tells Go to store all the arguments of Type in parameterName.
// The variadic variable can accept zero or more variables as the argument.

package main

import "fmt"

func numbers(i ...int) {
	fmt.Println(i)
}

func main() {
	numbers()
	numbers(9)
	numbers(99, 1000)
}
