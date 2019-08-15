package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//go:generate kgen ksvc
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, world")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(port, nil))
}
