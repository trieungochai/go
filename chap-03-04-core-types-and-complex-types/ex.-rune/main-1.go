package main

import "fmt"

func main() {
	username := "Sir_King_Über"

	for i := 0; i < len(username); i++ {
		fmt.Println(username[i])
	}
}

// go run main-1.go
// 83
// 105
// 114
// 95
// 75
// 105
// 110
// 103
// 95
// 195 --> Ü --> The numbers printed out are the byte values of the string. There are only 13 letters in our string. However, it contained a multi-byte character, so we printed out 14 byte values.
// 156 --> Ü --> The numbers printed out are the byte values of the string. There are only 13 letters in our string. However, it contained a multi-byte character, so we printed out 14 byte values.
// 98
// 101
// 114
