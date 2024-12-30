package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type messageData struct {
	Message string `json:"message"`
}

func getDataAndReturnResponse() messageData {
	// send the GET request
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	// get data from the response body
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// parse the response
	message := messageData{}
	err = json.Unmarshal(data, &message)
	if err != nil {
		log.Fatal(err)
	}

	// return the response data
	return message
}

func main() {
	data := getDataAndReturnResponse()
	fmt.Println(data.Message)
}
