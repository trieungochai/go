package main

import "fmt"

func swap(a *int, b *int) {
	// swap the values here
	*a, *b = *b, *a
}

func main() {
	a, b := 5, 10
	swap(&a, &b)

	fmt.Println(a == 10, b == 5)
}
