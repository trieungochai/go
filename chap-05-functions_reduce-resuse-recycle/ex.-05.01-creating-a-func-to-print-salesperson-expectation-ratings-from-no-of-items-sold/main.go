// create a function that will not have any parameters or return types.
// The function will iterate over a map and print the name and number of items sold on the map.
// It will also print a statement based on how the salesperson performed based on their sales.
package main

import (
	"fmt"
)

func main() {
	itemsSold()
}

func itemsSold() {
	items := make(map[string]int)
	items["John"] = 41
	items["Celina"] = 109
	items["Micah"] = 24

	for k, v := range items {
		fmt.Printf("%s sold %d items and ", k, v)
		if v < 40 {
			fmt.Println("is below expectations.")
		} else if v > 40 && v <= 100 {
			fmt.Println("meets expectations.")
		} else if v > 100 {
			fmt.Println("exceeded expectations.")
		}
	}
}
