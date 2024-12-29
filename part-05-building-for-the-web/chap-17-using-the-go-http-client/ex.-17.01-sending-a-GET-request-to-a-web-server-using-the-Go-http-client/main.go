// In this exercise,
// you will be getting data from a web server and printing out that data.
// You will send a GET request to https://www.google.com
// and display the data the web server returns
package main

import (
	"io"
	"log"
	"net/http"
)

func getDataAndReturnResponse() string {
	// use the default Go HTTP Client to request data from a server
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		log.Fatal(err)
	}

	// The data the server sends back is contained within r.Body,
	// so you just need to read that data.
	// To read the data within r.Body,
	// you can use the ReadAll function within the io package.
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// After have received the response from the server and read the data,
	// just need to return that data as a string
	return string(data)
}

func main() {
	data := getDataAndReturnResponse()
	log.Println(data)
}
