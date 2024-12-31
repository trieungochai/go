// Let’s imagine we want to make two calculations.
// First, we sum all numbers from 1 to 10, then numbers from 1 to 100.
// To save time, we want to make both calculations happen independently and see both results at the same time
package main

import (
	"log"
	"time"
)

// Create a function to sum 2 numbers
func sum(fromNum, toNum int) int {
	result := 0
	for i := fromNum; i <= toNum; i++ {
		result += i
	}

	return result
}

func main() {
	var batch1, batch2 int
	go func() {
		batch1 = sum(1, 100)
	}()
	batch2 = sum(1, 10)

	time.Sleep(time.Second)
	log.Println(batch1, batch2)
}
