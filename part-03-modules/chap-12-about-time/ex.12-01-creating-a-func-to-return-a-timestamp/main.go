// we will create a function called whatstheclock.
// The goal of this function is to demonstrate how you can create a function that wraps a nice,
// formatted time.Now() function and returns the date in an ANSIC format.
package main

import (
	"fmt"
	"time"
)

func whatstheclock() string {
	return time.Now().Format(time.ANSIC)
}

func main() {
	fmt.Println(whatstheclock())
}
