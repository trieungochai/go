// In this exercise, we will enhance our payDay() function so that it can recover from a panic.
// When our payDay() function panics, we will inspect the error from that panic.
// Then, depending on the error, we will print an informative message to the user.
package main

import (
	"errors"
	"fmt"
)

var (
	ErrHourlyRate  = errors.New("invalid hourly rate")
	ErrHoursWorked = errors.New("invalid hours worked")
)

func payDay(hoursWorked, hourlyRate int) int {
	defer func() {
		if r := recover(); r != nil {
			if r == ErrHourlyRate {
				fmt.Printf("hourly rate: %d\nerr: %v\n\n", hourlyRate, r)
			}
			if r == ErrHoursWorked {
				fmt.Printf("hours worked: %d\nerr: %v\n\n", hoursWorked, r)
			}
		}
		fmt.Printf("Pay was calculated based on:\nhours worked: %d\nhourly Rate: %d\n", hoursWorked, hourlyRate)
	}()

	if hourlyRate < 10 || hourlyRate > 75 {
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
	pay := payDay(100, 25)
	fmt.Println(pay)

	pay = payDay(100, 200)
	fmt.Println(pay)

	pay = payDay(60, 25)
	fmt.Println(pay)
}
