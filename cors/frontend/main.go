package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve static files from the "static" directory
	http.Handle("/", http.FileServer(http.Dir("static")))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
