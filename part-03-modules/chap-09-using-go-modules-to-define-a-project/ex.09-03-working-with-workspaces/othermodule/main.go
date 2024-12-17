package main

import (
	"fmt"

	"github.com/haitn/printer"
)

func main() {
	msg := printer.PrintNewUUID()
	fmt.Println(msg)
}
