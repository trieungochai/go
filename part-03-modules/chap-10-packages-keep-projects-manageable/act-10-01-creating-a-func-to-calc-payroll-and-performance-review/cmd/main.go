package main

import (
	"fmt"
	"os"

	"act.10.01/pkg/payroll"
)

var employeeReview = make(map[string]interface{})

func init() {
	fmt.Println("Welcome to the Employee Pay and Performance Review")
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++")
}

func init() {
	fmt.Printf("Initializing variables")
	employeeReview["WorkQuality"] = 5
	employeeReview["Teamwork"] = 2
	employeeReview["Communication"] = "Poor"
	employeeReview["Problem-solving"] = 4
	employeeReview["Dependability"] = "Unsatisfactory"
}

// create a main() func as an alias to the payroll package
func main() {
	d := payroll.Developer{Individual: payroll.Employee{Id: 1, FirstName: "Eric", LastName: "Davis"}, HourlyRate: 35, HoursWorkedInYear: 2400, Review: employeeReview}
	m := payroll.Manager{Individual: payroll.Employee{Id: 2, FirstName: "Mr.", LastName: "Boss"}, Salary: 150000, CommissionRate: .07}

	err := d.ReviewRating()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	payroll.PayDetails(d)
	payroll.PayDetails(m)
}
