// demonstrates the basics of using a panic
package main

import (
	"errors"
	"fmt"
)

func message(msg string) {
	if msg == "good-bye" {
		panic(errors.New("something went wrong"))
	}
}

func main() {
	thisMsg := "good-bye"
	message(thisMsg)
	fmt.Println("this line will not get printed")
}

// Code synopsis:
// - The function panics because the argument to the function message is "good-bye".
// - The panic() function prints the error message. Having a good error message helps with the debugging process.
// - Inside the panic, we are using errors.New(), which we used in the previous section to create an error type.
// - As you can see, fmt.Println() does not get executed in the main() function. Since there are no defer statements, execution stops immediately.
