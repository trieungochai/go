// using the make function,
// we’ll create several slices and display their length and capacity

package main

import "fmt"

// create a function that returns 3 int slices
func genSlices() ([]int, []int, []int) {
	// define a slice using the var notation
	var s1 []int
	// define a slice using make and set only the length
	s2 := make([]int, 10)
	// define a slice that uses both the length and capacity of the slices
	s3 := make([]int, 10, 50)

	return s1, s2, s3
}

func main() {
	s1, s2, s3 := genSlices()
	fmt.Printf("s1: len = %v cap = %v\n", len(s1), cap(s1))
	fmt.Printf("s2: len = %v cap = %v\n", len(s2), cap(s2))
	fmt.Printf("s3: len = %v cap = %v\n", len(s3), cap(s3))
}

// go run main.go
// s1: len = 0 cap = 0
// s2: len = 10 cap = 10
// s3: len = 10 cap = 50

// if you already know the maximum size your slice will need,
// setting the capacity upfront can improve performance
// because Go won’t have to spend extra resources resizing the underlying array
