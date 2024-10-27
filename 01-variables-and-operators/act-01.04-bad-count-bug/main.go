package main

import "fmt"

func main() {
	count := 0
	if count < 5 {
		count = 10 // Remove the ':=' to avoid shadowing
		count++
	}
	fmt.Println(count == 11)
}
