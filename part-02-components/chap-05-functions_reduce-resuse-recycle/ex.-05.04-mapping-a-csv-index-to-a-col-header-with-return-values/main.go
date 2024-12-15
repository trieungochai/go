// In Exercise 5.02 – mapping index values to column headers, we only printed the results of the index to the column header.
// In this exercise, we are going to return the map as the result. The map that’s being returned is the index-to-column header mapping.
// This is a more structured and reusable approach, especially if we want to use the resulting map in other parts of your code.
// By returning csvHeadersToColumnIndex from the csvHdrCol function, we can store or manipulate the result elsewhere in your program.
package main

import (
	"fmt"
	"strings"
)

func csvHdrCol(header []string) map[int]string {
	csvIdxToCol := make(map[int]string)
	for i, v := range header {
		v = strings.TrimSpace(v)
		switch strings.ToLower(v) {
		case "employee":
			csvIdxToCol[i] = v
		case "hours worked":
			csvIdxToCol[i] = v
		case "hourly rate":
			csvIdxToCol[i] = v
		}
	}
	return csvIdxToCol
}

func main() {
	hdr := []string{"empid", "employee", "address", "hours worked", "hourly rate", "manager"}
	result := csvHdrCol(hdr)
	fmt.Println("Result:")
	fmt.Println(result)
	fmt.Println()

	hdr2 := []string{"employee", "empid", "hours worked", "address", "manager", "hourly rate"}
	result2 := csvHdrCol(hdr2)
	fmt.Println("Result2:")
	fmt.Println(result2)
	fmt.Println()
}
