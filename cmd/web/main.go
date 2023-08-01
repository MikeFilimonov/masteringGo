package main

import (
	"fmt"
	"net/http"

	"github.com/MikeFilimonov/masteringGo/pkg/handlers"
)

const portNumber = ":8080"

// main is the main entry point of the app
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting the app at port %s", portNumber)

	_ = http.ListenAndServe(portNumber, nil)

}
