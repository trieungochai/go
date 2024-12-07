package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidLastName    = errors.New("invalid last name")
	ErrInvalidRoutingName = errors.New("invalid routing name")
)

func main() {
	fmt.Println(ErrInvalidLastName)
	fmt.Println(ErrInvalidRoutingName)
}
