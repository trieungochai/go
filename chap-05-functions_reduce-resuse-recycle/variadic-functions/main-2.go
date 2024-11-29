// The function takes the arguments being passed in and converts them into the new slice being specified.
package main

import "fmt"

func numbers(i ...int) {
	fmt.Println(i)
	fmt.Printf("%T\n", i)
	fmt.Printf("Len: %d\n", len(i))
	fmt.Printf("Cap: %d\n", cap(i))
}
func main() {
	numbers()
	numbers(1)
	numbers(99, 100)
}
