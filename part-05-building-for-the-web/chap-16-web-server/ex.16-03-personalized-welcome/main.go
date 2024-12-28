package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// the definition of a function that can be used as a handling function for an HTTP path
func Hello(w http.ResponseWriter, r *http.Request) {
	// save the query to a variable using the Query method URL from the request
	// The Query method on the URL object of the request returns a map[string][]string string with all the parameters sent through querystring in the URL.
	// We then assign this map to a variable, query.
	query := r.URL.Query()

	// at this point, we need to get the value of a specific parameter called name,
	// so we get the value from the name parameter
	// we have an assignment to 2 variables, but only one value comes from query["name"].
	// The 2d variable, ok, is a Boolean that tells us whether the name key exists.
	name, ok := query["name"]

	// If the name parameter has not been passed and we want an error message to appear,
	// we must add it if the variable is not found
	// in other words, if the ok variable is false
	// The conditional code gets called if the key does not exist in the slice,
	// and it writes a 400 code (bad request) to the header,
	// as well as a message to the response writer stating that the name has not been sent as a parameter.
	// We stop the execution with a return statement to prevent further actions.
	if !ok {
		w.WriteHeader(400)
		w.Write([]byte("Missing name"))
		return
	}

	// At this point, write a valid message to the response writer
	w.Write([]byte(fmt.Sprintf("Hello %s", strings.Join(name, ","))))
}

func main() {
	http.HandleFunc("/", Hello)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
