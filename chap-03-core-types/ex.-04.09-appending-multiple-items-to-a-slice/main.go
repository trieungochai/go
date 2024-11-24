// we’ll use the variadic parameter of append to add multiple values in the form of predefined data to a slice.
// Then, we’ll add a dynamic amount of data based on user input to the same slice.
package main

import (
	"fmt"
	"os"
)

func getPassedArgs() []string {
	// retrieves command-line arguments passed to the program,
	// starting from index 1 (since os.Args[0] is the program name)
	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}

	return args
}

// accepts a slice of additional locales (extraLocales) as an argument
func getLocales(extraLocales []string) []string {
	var locales []string

	// starts with a default set of locales: en_US and fr_FR
	locales = append(locales, "en_US", "fr_FR")
	// uses the variadic ... operator to append extraLocales to the default locales slice
	locales = append(locales, extraLocales...)

	return locales
}

func main() {
	// retrieves additional locales
	// combines the default locales with the additional ones
	locales := getLocales(getPassedArgs())
	fmt.Println("Locales to use:", locales)
}
