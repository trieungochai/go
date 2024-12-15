// We will be implementing a program that will calculate the area of a triangle, rectangle, and square.
// The program will use a single function that accepts a Shape interface.
// Any type that satisfies the Shape interface can be passed as an argument to the function.
// This function should then print the area and the name of the shape.

package main

import "fmt"

type Shape interface {
	Area() float64
	Name() string
}

type triangle struct {
	base   float64
	height float64
}

type rectangle struct {
	length float64
	width  float64
}

type square struct {
	side float64
}

func main() {
	t := triangle{base: 15.5, height: 20.1}
	r := rectangle{length: 20, width: 10}
	s := square{side: 10}

	printShapeDetails(t, r, s)
}

func (t triangle) Area() float64 {
	return (t.base * t.height) / 2
}
func (t triangle) Name() string {
	return "triangles"
}

func (r rectangle) Area() float64 {
	return r.length * r.width
}
func (r rectangle) Name() string {
	return "rectangle"
}

func (s square) Area() float64 {
	return s.side * s.side
}
func (s square) Name() string {
	return "square"
}

func printShapeDetails(shapes ...Shape) {
	for _, item := range shapes {
		fmt.Printf("The area of %s is: %.2f\n", item.Name(), item.Area())
	}
}
