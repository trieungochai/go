// going to read from a map using direct access and a loop
// also check to see if a key exists in the map
package main

import (
	"fmt"
	"os"
)

// User Data
// returns a hard-coded map of user IDs and their corresponding names
func getUsers() map[string]string {
	return map[string]string{
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
		"073": "Tracy",
	}
}

// User Retrieval
// checks if a given user ID exists in the user map
// and returns the corresponding name and a boolean flag indicating existence
func getUser(id string) (string, bool) {
	users := getUsers()
	user, exists := users[id]
	return user, exists
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("User ID not passed")
		os.Exit(1)
	}

	userId := os.Args[1]
	name, exists := getUser(userId)
	if !exists {
		fmt.Printf("Passed user ID (%v) not found.\nUsers:\n", userId)
		for key, value := range getUsers() {
			fmt.Println("    ID:", key, "Name:", value)
		}
		os.Exit(1)
	}
	fmt.Println("Name:", name)
}
