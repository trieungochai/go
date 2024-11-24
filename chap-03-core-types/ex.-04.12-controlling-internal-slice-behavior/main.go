// going to explore 5 different ways to copy data from slice to slice
// and how that has an impact on a slice’s internal behavior
package main

import "fmt"

func linked() (int, int, int) {
	// define an int slice, initialized with some data
	s1 := []int{1, 2, 3, 4, 5}
	// then, make a simple variable copy of that slice
	s2 := s1
	// create a new slice by copying all the values from the first slice as part of a slice range operation
	s3 := s1[:]
	// change some data in the first slice. Later, we’ll see how this affects the second and third slices
	s1[3] = 99
	return s1[3], s2[3], s3[3]
}

func noLink() (int, int) {
	// define a slice with some data and do a simple copy again
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	// this time, we’ll append to the first slice before we do anything else
	// this operation changes the length and capacity of the slice
	s1 = append(s1, 6)
	s1[3] = 99
	return s1[3], s2[3]
}

func capLinked() (int, int) {
	s1 := make([]int, 5, 10)
	// fill the first array with the same data as before
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1
	// append a new value to the first slice, which changes its length but not its capacity
	s1 = append(s1, 6)
	s1[3] = 99
	return s1[3], s2[3]
}

// use make again to set a capacity,
// but use append to add elements that will go beyond that capacity
func capNoLink() (int, int) {
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1
	s1 = append(s1, []int{10: 11}...)
	s1[3] = 99
	return s1[3], s2[3]
}

// use copy to copy the elements from the first slice to the second slice.
// copy returns how many elements were copied from one slice to another
func copyNoLink() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, len(s1))
	copied := copy(s2, s1)
	s1[3] = 99
	return s1[3], s2[3], copied
}

// use append to copy the value into the second slice
// using append in this way results in the values being copied into a new hidden array
func appendNoLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := append([]int{}, s1...)
	s1[3] = 99
	return s1[3], s2[3]
}

func main() {
	l1, l2, l3 := linked()
	fmt.Println("Linked        :", l1, l2, l3)

	nl1, nl2 := noLink()
	fmt.Println("No Link       :", nl1, nl2)

	cl1, cl2 := capLinked()
	fmt.Println("Cap Link      :", cl1, cl2)

	cnl1, cnl2 := capNoLink()
	fmt.Println("Cap No Link   :", cnl1, cnl2)

	copynl1, copynl2, copied := copyNoLink()
	fmt.Println("Copy No Link  :", copynl1, copynl2, copied)

	anl1, anl2 := appendNoLink()
	fmt.Println("Append No Link:", anl1, anl2)
}
