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
	http.Handle("/", http.FileServer(http.Dir("./dist")))

	log.Printf("Server listening on port %d", port)

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
