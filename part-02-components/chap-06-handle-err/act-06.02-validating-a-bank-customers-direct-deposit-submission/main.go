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
	ErrInvalidRoutingNumber = errors.New("invalid routing number")
)

func (dd *directDeposit) validateRoutingNumber() error {
	if dd.routingNumber < 100 {
		return ErrInvalidRoutingNumber
	}

	return nil
}

func (dd *directDeposit) validateLastName() error {
	dd.lastName = strings.TrimSpace(dd.lastName)
	if len(dd.lastName) == 0 {
		return ErrInvalidLastName
	}

	return nil
}

func (dd *directDeposit) report() {
	fmt.Println(strings.Repeat("*", 30))
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
		bankName:      "XYZ Inc",
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
