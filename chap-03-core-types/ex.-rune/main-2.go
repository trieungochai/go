package main

import "fmt"

func main() {
	username := "Sir_King_Über"

	for i := 0; i < len(username); i++ {
		fmt.Print(string(username[i]))
	}
}

// Sir_King_Ãber
// The output is as expected until we get to the Ü character.
// That’s because Ü was encoded using more than one byte, and each byte on its own no longer makes sense.
// To safely work with individual characters of a multi-byte string, you first must convert the string slice of byte types to a slice of rune types.
