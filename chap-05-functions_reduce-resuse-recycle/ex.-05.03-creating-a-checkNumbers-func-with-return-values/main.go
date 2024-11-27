// it accepts only an integer.
// leave the onus on the caller to perform the looping if they desire to do so.
// going to have two returns.
// The first will be the number provided and the corresponding text indicating if the number is Even or Odd

package main

import "fmt"

// simplify the if{}else{} statements by replacing them with switch statements.
// The return statement will immediately stop the execution of the function and return the results to the caller.
func checkNumbers(value int) (int, string) {
	switch {
	case value%2 == 0:
		return value, "EVEN"
	default:
		return value, "ODD"
	}
}

func main() {
	// assign variables to the return values of our function.
	// The n, and s variables correspond to the values being returned from our function, which are int and string
	for i := 0; i <= 15; i++ {
		value, result := checkNumbers(i)
		fmt.Printf("Results:  %d %s\n", value, result)
	}
}
