// going to sum up a variable number of arguments.
// pass the arguments as a list of arguments and as a slice.
// The return value will be an int type – that is, the sum of the values we passed to the function.

package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	i := []int{5, 10, 15}
	fmt.Println(sum(5, 4))
	fmt.Println(sum(i...))
}
