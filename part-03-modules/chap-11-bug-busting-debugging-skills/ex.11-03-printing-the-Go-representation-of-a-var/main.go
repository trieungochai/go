package main

import "fmt"

type person struct {
	lastName string
	age      int
	salary   float64
}

func main() {
	firstName := "Kai"
	grades := []int{100, 87, 67}
	states := map[string]string{"KY": "Kentucky", "HCM": "Ho Chi Minh", "HN": "Hanoi"}
	p := person{lastName: "Cai", age: 69, salary: 1000}

	fmt.Printf("firstName value %#v\n", firstName)
	fmt.Printf("firstName value %T\n", firstName)
	fmt.Printf("grades value %#v\n", grades)
	fmt.Printf("grades value %T\n", grades)
	fmt.Printf("states value %#v\n", states)
	fmt.Printf("states value %T\n", states)
	fmt.Printf("p value %#v\n", p)
	fmt.Printf("p value %T\n", p)
}
