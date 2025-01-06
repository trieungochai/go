// In this exercise,
// we want to have one Goroutine to send messages and another one to print them.
// Moreover, we want to know when the sender has finished sending messages.

package main

import "log"

// define a function that will first receive the strings and print them later
func readThem(in, out chan string) {
	// Then, create a loop over the channel until the channel is closed
	for i := range in {
		log.Println(i)
	}

	// Finally, send a notification saying that the processing has finished
	out <- "done"
}

func main() {
	// set the log flags to 0 so that we do not see anything other than the strings we send
	log.SetFlags(0)

	// create the necessary channels and use them to spin up the Goroutine
	in, out := make(chan string), make(chan string)
	go readThem(in, out)

	// create a set of strings and loop over them, sending each string to the channel
	strs := []string{"a", "b", "c", "d", "e", "f"}
	for _, s := range strs {
		in <- s
	}

	// close the channel you used to send the messages and wait for the done signal
	close(in)
	<-out
}

// --------
// We see that the main() function has received all the messages from the Goroutine and has printed them.
// The main() function terminates only when it has been notified that all incoming messages have been sent.
