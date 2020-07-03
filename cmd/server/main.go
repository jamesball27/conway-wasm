package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port = 3000
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})

	log.Printf("Server listening on port %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
