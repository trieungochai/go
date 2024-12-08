// The bank has now decided that it would rather crash the program when an invalid routing number is submitted.
// The bank feels that the erroneous data should cause the program to stop processing the direct deposit data.
// Build this on top of Activity 6.02 – validating a bank customer’s direct deposit submission.
package main

import (
	"errors"
	"fmt"
	"strings"
)

type directDeposit struct {
	lastName      string
	firstName     string
	bankName      string
	routingNumber int
	accountNumber int
}

var (
	ErrInvalidLastName      = errors.New("invalid last name")
	ErrInvalidRoutingNumber = errors.New("invaliad routing number")
)

func (dd *directDeposit) validateLastName() error {
	dd.lastName = strings.TrimSpace(dd.lastName)
	if len(dd.lastName) == 0 {
		return ErrInvalidLastName
	}

	return nil
}

func (dd *directDeposit) validateRoutingNumber() error {
	if dd.routingNumber < 100 {
		panic(ErrInvalidRoutingNumber)
	}

	return nil
}

func (dd *directDeposit) report() {
	fmt.Println(strings.Repeat("*", 80))
	fmt.Println("Last Name: ", dd.lastName)
	fmt.Println("First Name: ", dd.firstName)
	fmt.Println("Bank Name: ", dd.bankName)
	fmt.Println("Routing Number: ", dd.routingNumber)
	fmt.Println("Account Number: ", dd.accountNumber)
}

func main() {
	dd := directDeposit{
		lastName:      "  ",
		firstName:     "Abe",
		bankName:      "WilkesBooth Inc",
		routingNumber: 17,
		accountNumber: 1809,
	}

	err := dd.validateRoutingNumber()
	if err != nil {
		fmt.Println(err)
	}

	err = dd.validateLastName()
	if err != nil {
		fmt.Println(err)
	}

	dd.report()
}
