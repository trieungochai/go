// In this exercise, we’ll declare a string and initialize it with a multi-byte string value.
// We’ll then loop over the string using range to give us each character, one at a time.
// We’ll then print out the byte index and the character to the console.

package main

import "fmt"

func main() {
	logLevel := "デバッグ"
	for index, runeVal := range logLevel {
		fmt.Println(index, string(runeVal))
	}
}
