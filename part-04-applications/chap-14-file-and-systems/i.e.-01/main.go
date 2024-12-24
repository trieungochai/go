package main

import (
	"flag"
	"fmt"
)

func main() {
	// v variable will reference the value for either -value or --value
	// initial value of *v is the default value of -1 before calling flag.Parse()
	v := flag.Int("value", -1, "Need a value for the flag.")

	// after defining the flags, you must call flag.Parse() to parse the defined flags into the command line
	// calling flag.Parse() places the argument for -value into *v
	flag.Parse()

	// once you have called the flag.Parse() function, the flags will be available
	fmt.Println(*v)
}

// On the command line, execute the following command and you will get the executable in the same directory:
// go build -o flagapp main.go
// To get the executable on Windows, run:
// go build -o flagapp.exe main.go
