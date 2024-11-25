package main

import "fmt"

func salesTaxCalculator(cost float64, taxRate float64) float64 {
	return cost * taxRate
}

func main() {
	taxTotal := .0
	// Cake
	taxTotal += salesTaxCalculator(.99, .075)
	// Milk
	taxTotal += salesTaxCalculator(2.75, .015)
	// Butter
	taxTotal += salesTaxCalculator(.87, .02)
	// Total
	fmt.Println("Sales Tax Total: ", taxTotal)
}

// go run main.go
// Sales Tax Total:  0.1329
