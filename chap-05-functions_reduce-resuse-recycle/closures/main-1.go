package main

import "fmt"

func main() {
	// initialize a variable in the main() function called i and set it to 0
	i := 0

	// assign incrementor to our anonymous function
	// The anonymous function increments i and returns it.
	// Notice that our function does not have any input parameters.
	incrementor := func() int {
		i += 1
		return i
	}

	fmt.Println(incrementor())
	fmt.Println(incrementor())
	i += 10
	fmt.Println(incrementor())
}

// Notice that, outside our function, we increment i by 10.
// This is a problem. We want i to be isolated and for it not to change as this is not the desired behavior.
// When we print the results of incrementor again, it will be 12. We want it to be 3.
// We will correct this in our next example.
