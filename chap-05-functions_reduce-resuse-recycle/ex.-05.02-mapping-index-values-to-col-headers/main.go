// function that we are going to create will be taking a slice of column headers from a CSV file.
// It will print out a map of an index value of the headers we are interested in

package main

import (
	"fmt"
	"strings"
)

// filters and maps specific header values to their corresponding column indices in the header slice
func csvHdrCol(header []string) {
	csvHeadersToColumnIndex := make(map[int]string)
	for i, v := range header {
		v = strings.TrimSpace(v)
		switch strings.ToLower(v) {
		case "employee":
			csvHeadersToColumnIndex[i] = v
		case "hours worked":
			csvHeadersToColumnIndex[i] = v
		case "hourly rate":
			csvHeadersToColumnIndex[i] = v
		}
	}
	fmt.Println(csvHeadersToColumnIndex)
}

func main() {
	hdr1 := []string{"empid", "employee", "address", "hours worked", "hourly rate", "manager"}
	csvHdrCol(hdr1)
	hdr2 := []string{"employee", "empid", "hours worked", "address", "manager", "hourly rate"}
	csvHdrCol(hdr2)
}
