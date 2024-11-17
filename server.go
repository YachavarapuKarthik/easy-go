package main

import (
	"fmt"
	"net/http"
)

// Handler function for the root route
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to your Go web server!")
}

func main2() {
	// Register the handler function for the root route
	http.HandleFunc("/", handler)

	// Start the server on port 8080
	fmt.Println("Starting server at port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
