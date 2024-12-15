package main

import (
	"fmt"
	"os"
)

var users = map[string]string{
	"305": "Sue",
	"204": "Bob",
	"631": "Jake",
	"073": "Tracy",
}

func getUsername(id string) (string, bool) {
	username, exists := users[id]
	return username, exists
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("User ID not passed")
		os.Exit(1)
	}

	username, exists := getUsername(os.Args[1])
	if !exists {
		fmt.Printf("error: user (%v) not found", os.Args[1])
		os.Exit(1)
	}

	fmt.Println("Hi,", username)
}
