// One problem with the previous example that we noticed is that any code in the main function has access to i.
// i can be accessed and changed outside of our function.
// This is not the desired behavior;
// we want the incrementor to be the only one to change that value.
// In other words, we want i to be protected from other functions changing it.
// The only function that should be changing is our anonymous function when we call it
package main

import "fmt"

// declared a function called incrementor(). This function has a return type of func() int.
func incrementor() func() int {
	// Using i := 0, we initialize our variable at the level of the incrementor() function;
	// this is similar to what we did in the previous example,
	// except it was at the main() function level and anyone at that level had access to i.
	// Only the incrementor() function has access to the i variable with this implementation.
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	// In the main() function, increment := incrementor() assigns a variable to func() int that gets returned.
	// It is important to note that incrementor() only gets executed once here. In our main() function,
	// it is no longer being referenced or executed.
	increment := incrementor()

	// increment() is of the func() int type. Each call to increment() runs the anonymous function code.
	// It is referencing the i variable, even after incrementor() has been executed.
	fmt.Println(increment())
	fmt.Println(increment())
}

// the preceding example demonstrated how we can protect our variable by wrapping it with an anonymous function,
// thereby restricting access to updating the variable only through invoking the anonymous function itself.
// This is shown through the expected output, where we’ve incremented i twice: 1, 2
