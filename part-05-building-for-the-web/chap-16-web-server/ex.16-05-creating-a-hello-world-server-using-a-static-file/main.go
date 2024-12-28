// In this exercise, you learned how to use a static HTML file to serve a web page,
// as well as how detaching the static resources from your app
// allows you to change your served page without having to restart your app.
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
