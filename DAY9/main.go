package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

// Response struct
type Response struct {
    Message string `json:"message"`
}

// Handler for GET requests
func helloHandler(w http.ResponseWriter, r *http.Request) {
    response := Response{Message: "Hello, World!"}
    json.NewEncoder(w).Encode(response)
}

// Handler for POST requests
func echoHandler(w http.ResponseWriter, r *http.Request) {
    var requestBody map[string]string
    json.NewDecoder(r.Body).Decode(&requestBody)
    json.NewEncoder(w).Encode(requestBody)
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/echo", echoHandler)

    fmt.Println("Server running on port 8080...")
    http.ListenAndServe(":8080", nil)
}
