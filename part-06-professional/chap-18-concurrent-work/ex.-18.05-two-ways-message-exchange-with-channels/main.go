package main

import (
	"fmt"
	"log"
)

func greet(ch chan string) {
	receivedMsg := <-ch
	ch <- fmt.Sprintf("Thanks for %s", receivedMsg)
	ch <- "Hello David"
}

func main() {
	ch := make(chan string)
	go greet(ch)

	ch <- "Hello John"
	log.Println(<-ch)
	log.Println(<-ch)
}
