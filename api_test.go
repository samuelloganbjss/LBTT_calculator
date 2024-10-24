package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculateHandler(t *testing.T) {
	// Set up a request with a valid price
	reqBody := RequestBody{Price: 100000}
	body, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", "/calculate", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(calculateHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	var resBody ResponseBody
	err = json.NewDecoder(rr.Body).Decode(&resBody)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	// Validate the LBTT value (assuming a 5% calculation in this case)
	expectedLbtt := 5000.0
	if resBody.Lbtt != expectedLbtt {
		t.Errorf("handler returned unexpected LBTT: got %v want %v", resBody.Lbtt, expectedLbtt)
	}
}

func TestInvalidPrice(t *testing.T) {
	// Test with an invalid price (negative value)
	reqBody := RequestBody{Price: -100000}
	body, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", "/calculate", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(calculateHandler)

	handler.ServeHTTP(rr, req)

	// Check for bad request response
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code for invalid price: got %v want %v",
			status, http.StatusBadRequest)
	}
}
