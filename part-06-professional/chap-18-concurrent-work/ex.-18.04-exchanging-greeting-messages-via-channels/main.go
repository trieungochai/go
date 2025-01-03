// we will use a Goroutine to send a greeting message,
// and then we will receive the greeting in the main process.
package main

import "log"

func greet(ch chan string) {
	ch <- "See you Space Cowboy!!!"
}

func main() {
	// instantiate a channel and pass it to the greeter() function
	// Here, only a channel of strings is created and passed as a parameter to the call to a new routine called greet.
	ch := make(chan string)
	go greet(ch)

	log.Println(<-ch)
}
