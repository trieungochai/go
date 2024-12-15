package main

import (
	"fmt"
	"unicode"
)

func passwordChecker(pw string) bool {
	// make it UTF-8 safes
	pwR := []rune(pw)

	if len(pwR) < 8 {
		return false
	}
	if len(pwR) > 15 {
		return false
	}

	hasLowerCaseLetter := false
	hasUpperCaseLetter := false
	hasNumber := false
	hasSymbol := false

	// loop over the multi-byte characters one at a time
	for _, v := range pwR {
		if unicode.IsLower(v) {
			hasLowerCaseLetter = true
		}
		if unicode.IsUpper(v) {
			hasUpperCaseLetter = true
		}
		if unicode.IsNumber(v) {
			hasNumber = true
		}
		if unicode.IsSymbol(v) || unicode.IsPunct(v) {
			hasSymbol = true
		}
	}

	return hasLowerCaseLetter && hasUpperCaseLetter && hasNumber && hasSymbol
}

func main() {
	if passwordChecker("") {
		fmt.Println("password good")
	} else {
		fmt.Println("password bad")
	}
	if passwordChecker("This!I5A") {
		fmt.Println("password good")
	} else {
		fmt.Println("password bad")
	}
}
