package main

import (
	"fmt"
	"strings"
)

func main() {
	var str interface{} = "some string"
	v := str.(string)
	fmt.Println(strings.Title(v))
}

// The preceding code asserts that `str` is of the string type and assigns it to the variable v
// Since v is a string, it will print it with title casing
// The result is as follows: Some String
