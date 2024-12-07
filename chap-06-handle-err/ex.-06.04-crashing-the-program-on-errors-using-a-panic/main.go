package main

import (
	"errors"
	"fmt"
)

var (
	ErrHourlyRate  = errors.New("invalid hourly rate")
	ErrHoursWorked = errors.New("invalid hours worked")
)

// Inside the payDay() function, assign a variable, report, to an anonymous function.
// This anonymous function provides details of the arguments provided to the payDay() function.
// Even though we are not returning errors, this will provide some insight as to why the function panics.
// Since it is a deferred function, it will always execute before the function exits
func payDay(hoursWorked, hourlyRate int) int {
	// When a panic occurs in the payDay() function, the defer function, report(),
	// will give the caller some insight into why the panic occurred.
	// The panic will bubble up the stack to the main() function and execution will stop immediately.
	report := func() {
		fmt.Printf("HoursWorked: %d\nHourldyRate: %d\n", hoursWorked, hourlyRate)
	}
	defer report()

	if hourlyRate < 40 || hourlyRate > 75 {
		panic(ErrHourlyRate)
	}

	if hoursWorked < 0 || hoursWorked > 80 {
		panic(ErrHoursWorked)
	}

	if hoursWorked > 40 {
		hoursOver := hoursWorked - 40
		overTime := hoursOver * 2
		regularPay := hoursWorked * hourlyRate
		return regularPay + overTime
	}

	return hoursWorked * hourlyRate
}

func main() {
	pay := payDay(81, 50)
	fmt.Println(pay)
}
