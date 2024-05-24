// greet_service.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

const helloServiceURL = "http://localhost:8001/hello"

type NameResponse struct {
	Name   string `json:"name"`
	Length int    `json:"length"`
}

func GetLastStringFromURL(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) > 2 {
		return parts[2]
	}
	return ""
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := GetLastStringFromURL(r.URL.Path)
	if name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	response, err := http.Post(helloServiceURL, "application/json", strings.NewReader(fmt.Sprintf(`{"name": "%s"}`, name)))
	if err != nil || response.StatusCode != http.StatusOK {
		http.Error(w, "Failed to call hello service", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	var nameResponse NameResponse
	err = json.NewDecoder(response.Body).Decode(&nameResponse)
	if err != nil {
		http.Error(w, "Failed to parse hello service response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nameResponse)
}

func main() {
	http.HandleFunc("/greet/", greetHandler)
	port := "8000"
	fmt.Printf("Starting greet service on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
