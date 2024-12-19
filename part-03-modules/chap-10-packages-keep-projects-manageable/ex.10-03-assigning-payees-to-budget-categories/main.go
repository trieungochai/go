// We are going to expand our program from Exercise 10.02.
// Loading budget categories, to now assign payees to budget categories.
// This is similar to many budgeting applications that try to match payees to commonly used categories.
// We will then print the mapping of a payee to a category.
package main

import "fmt"

var budgetCategories = make(map[int]string)
var payeeToCategories = make(map[string]int)

func init() {
	fmt.Println("Initializing our budgetCategories")
	budgetCategories[1] = "Car Insurance"
	budgetCategories[2] = "Mortgage"
	budgetCategories[3] = "Electricity"
	budgetCategories[4] = "Retirement"
	budgetCategories[5] = "Vacation"
	budgetCategories[7] = "Groceries"
	budgetCategories[8] = "Car Payment"
}

func init() {
	fmt.Println("Assign our Payees to Categories")
	payeeToCategories["Nationwide"] = 1
	payeeToCategories["BBT Loan"] = 2
	payeeToCategories["1st Energy Electric"] = 3
	payeeToCategories["Ameriprise Financial"] = 4
	payeeToCategories["Walt Disney World"] = 5
	payeeToCategories["ALDI"] = 7
	payeeToCategories["Martins"] = 7
	payeeToCategories["Wal Mart"] = 7
	payeeToCategories["Chevy Loan"] = 8
}

func main() {
	for k, v := range payeeToCategories {
		fmt.Printf("Payee: %s, Category: %s\n", k, budgetCategories[v])
	}
}
