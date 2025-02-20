package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHelloHandler tests the GET /hello endpoint
func TestHelloHandler(t *testing.T) {
	// Create a new request
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Create the handler
	handler := http.HandlerFunc(helloHandler)

	// Serve the request
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	expected := Response{Message: "Hello, World!"}
	var got Response
	err = json.NewDecoder(rr.Body).Decode(&got)
	if err != nil {
		t.Fatal(err)
	}

	if got.Message != expected.Message {
		t.Errorf("handler returned unexpected body: got %v want %v",
			got.Message, expected.Message)
	}
}

// TestEchoHandler tests the POST /echo endpoint
func TestEchoHandler(t *testing.T) {
	// Create test data
	testData := map[string]string{
		"message": "Test Message",
		"name":    "Test Name",
	}

	// Convert test data to JSON
	jsonData, err := json.Marshal(testData)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new request with the JSON body
	req, err := http.NewRequest("POST", "/echo", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Create the handler
	handler := http.HandlerFunc(echoHandler)

	// Serve the request
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	var got map[string]string
	err = json.NewDecoder(rr.Body).Decode(&got)
	if err != nil {
		t.Fatal(err)
	}

	// Compare the response with the test data
	for key, value := range testData {
		if got[key] != value {
			t.Errorf("handler returned unexpected value for key %s: got %v want %v",
				key, got[key], value)
		}
	}
}
