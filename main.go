package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// main is the main entry point of the app
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting the app at port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)

}
