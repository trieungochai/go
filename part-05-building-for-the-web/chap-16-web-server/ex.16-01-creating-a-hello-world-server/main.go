package main

import (
	"log"
	"net/http"
)

// create handler, the struct that will handle the requests
type hello struct{}

func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>See you! Space Cowboy</h1>"
	w.Write([]byte(msg))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", hello{}))
}
