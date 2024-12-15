package main

import "fmt"

// create a struct called record that will store the key, type of value, and data from map[string]interface{}.
// This struct is used to store the analysis that we are performing on the map.
// The key field is the name of the map key.
// The valueType field stores the type of data stored as a value in the map.
// The data field stores the data we are analyzing.
// It is an empty interface{}, since there can be various types of data in the map.
type record struct {
	key       string
	valueType string
	data      interface{}
}

// create a person struct that will be added to our map[string]interface{}
type person struct {
	lastName  string
	age       int
	isMarried bool
}

// create an animal struct that will be added to our map[string]interface{}
type animal struct {
	name     string
	category string
}

func main() {
	// initialize our map.
	// The map is initialized to a string for the key and an empty interface for the value.
	// We then assign a to an animal struct literal and p to a person struct literal.
	// Then, we start adding various key-value pairs to the map
	m := make(map[string]interface{})
	a := animal{name: "oreo", category: "cat"}
	p := person{lastName: "Doe", isMarried: false, age: 19}

	m["person"] = p
	m["animal"] = a
	m["age"] = 54
	m["isMarried"] = true
	m["lastName"] = "Smith"

	// Next, we initialize a slice of record. We iterate over the map and add records to rs:
	rs := []record{}
	for k, v := range m {
		r := newRecord(k, v)
		rs = append(rs, r)
	}

	for _, v := range rs {
		fmt.Println("Key: ", v.key)
		fmt.Println("Data: ", v.data)
		fmt.Println("Type: ", v.valueType)
		fmt.Println()
	}
}

// create a newRecord() function. The key parameter will be our map’s key.
// The function also takes interface{} as an input parameter.
// i will be our map’s value for the key that is passed to the function.
// It will return a record type.
func newRecord(key string, i interface{}) record {
	// Inside the newRecord() function, we initialize record{} and assign it to the r variable.
	// We then assign r.key to the key input parameter.
	r := record{}
	r.key = key

	// The switch statement assigns the type of i to the v variable.
	// The v variable type gets evaluated against a series of case statements.
	// If a type evaluates to true for one of the case statements,
	// then the valueType record gets assigned to that type, along with the value of v to r.data, and then returns the record type.
	switch v := i.(type) {
	case int:
		r.valueType = "int"
		r.data = v
	case bool:
		r.valueType = "bool"
		r.data = v
	case string:
		r.valueType = "string"
		r.data = v
	case person:
		r.valueType = "person"
		r.data = v
	default:
		r.valueType = "unknown"
		r.data = v
	}
	return r
}
