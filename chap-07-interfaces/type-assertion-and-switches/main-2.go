// what will happen if s is not of type T?
// Having a panic being thrown is not desirable.
// However, Go has a way to check whether str is a string.
package main

import (
	"fmt"
)

func main() {
	var str interface{} = "some string"
	v, isValid := str.(int)
	fmt.Println(v, isValid)
}

// 1. A type assertion returns two values, the underlying value and a Boolean value.

// 2. isValid is assigned to a return type of bool.
// If it returns true, that indicates that str is of the int type.
// It means that the assertion is true.
// We can use the Boolean that was returned to determine what action we can take on str.

// 3. When the assertion fails, it will return false.
// The return value will be the zero value that you are trying to assert to.
// It also will not panic.
