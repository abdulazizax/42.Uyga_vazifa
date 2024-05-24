// hello_service.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type NameRequest struct {
	Name string `json:"name"`
}

type NameResponse struct {
	Name   string `json:"name"`
	Length int    `json:"length"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var nameRequest NameRequest
	err := json.NewDecoder(r.Body).Decode(&nameRequest)
	if err != nil || nameRequest.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	nameResponse := NameResponse{Name: nameRequest.Name, Length: len(nameRequest.Name)}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nameResponse)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	port := "8001"
	fmt.Printf("Starting hello service on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
