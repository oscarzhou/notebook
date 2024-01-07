package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Cat struct {
	Name   string
	Age    int
	Breed  string
	Color  string
	Weight float64
}

func init() {
	dummyCats = make([]Cat, 5)
	dummyCats[0] = Cat{Name: "Whiskers", Age: 3, Breed: "Siamese", Color: "Brown", Weight: 8.5}
	dummyCats[1] = Cat{Name: "Mittens", Age: 2, Breed: "Tabby", Color: "Gray", Weight: 7.2}
	dummyCats[2] = Cat{Name: "Fluffy", Age: 4, Breed: "Persian", Color: "White", Weight: 10.0}
	dummyCats[3] = Cat{Name: "Tiger", Age: 5, Breed: "Bengal", Color: "Spotted", Weight: 12.3}
	dummyCats[4] = Cat{Name: "Smokey", Age: 1, Breed: "Russian Blue", Color: "Blue", Weight: 6.8}

}

var dummyCats []Cat

func main() {
	http.HandleFunc("/cats", listCatsHandler)
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func listCatsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080") // Allow CORS from frontend origin
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	// Handle preflight requests
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	resp, err := json.Marshal(dummyCats)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s", resp)
}
