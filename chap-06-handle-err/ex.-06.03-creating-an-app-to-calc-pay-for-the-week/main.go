// we are going to create a function that calculates pay for the week.
// This function will accept two arguments – the hours worked during the week and the hourly rate.
// The function is going to check whether the two parameters meet the criteria for being valid.
// The function will need to calculate regular pay, which is hours less than or equal to 40, and overtime pay, which is hours greater than 40 for the week.
// We will create two error values using errors.New().
// The one error value will be used when there is an invalid hourly rate.
// An invalid hourly rate in our app is an hourly rate that is less than 10 or greater than 75.
// The second error value will be when hours per week are not between 0 and 80.

package main

import (
	"errors"
	"fmt"
)

var (
	ErrHourlyRate  = errors.New("invalid hourly rate")
	ErrHoursWorked = errors.New("invalid hours worked")
)

func payDay(hoursWorked, hourlyRate int) (int, error) {
	if hourlyRate < 40 || hourlyRate > 75 {
		return 0, ErrHourlyRate
	}

	if hoursWorked < 0 || hoursWorked > 80 {
		return 0, ErrHoursWorked
	}

	if hoursWorked > 40 {
		hoursOver := hoursWorked - 40
		overTime := hoursOver * 2
		regularPay := hoursWorked * hourlyRate
		return regularPay + overTime, nil
	}

	return hoursWorked * hourlyRate, nil
}

func main() {
	pay, err := payDay(81, 50)
	if err != nil {
		fmt.Println(err)
	}

	pay, err = payDay(80, 5)
	if err != nil {
		fmt.Println(err)
	}

	pay, err = payDay(80, 50)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pay)
}
