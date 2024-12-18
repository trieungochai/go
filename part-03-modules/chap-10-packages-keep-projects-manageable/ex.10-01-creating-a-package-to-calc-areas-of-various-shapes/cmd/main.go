package main

import "ex.10.01/pkg/shape"

func main() {
	t := shape.Triangle{Base: 15.5, Height: 20.1}
	r := shape.Rectangle{Length: 20, Width: 10}
	s := shape.Square{Side: 10}

	shape.PrintShapeDetails(t, r, s)
}

// Type the following:
// go build
// The go build command will compile your program and create an executable named after the directory, cmd.
// Type the executable name and hit Enter:
// ./cmd
// The expected output is as follows:
// The area of Triangle is: 155.78
// The area of Rectangle is: 200.00
// The area of Square is 100.00
