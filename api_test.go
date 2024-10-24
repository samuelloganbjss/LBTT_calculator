package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculateLBTTHandler_StandardBuyer(t *testing.T) {
	reqBody := RequestBody{Price: 300000, IsFirstTimeBuyer: false, IsAdditionalDwelling: false}
	body, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", "/calculate", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(calculateLBTTHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var resBody ResponseBody
	err = json.NewDecoder(rr.Body).Decode(&resBody)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	expectedLbtt := 4600.0
	if resBody.Lbtt != expectedLbtt {
		t.Errorf("handler returned unexpected LBTT: got %v want %v", resBody.Lbtt, expectedLbtt)
	}
}

func TestCalculateLBTTHandler_FirstTimeBuyer(t *testing.T) {
	reqBody := RequestBody{Price: 300000, IsFirstTimeBuyer: true, IsAdditionalDwelling: false}
	body, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", "/calculate", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(calculateLBTTHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var resBody ResponseBody
	err = json.NewDecoder(rr.Body).Decode(&resBody)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	expectedLbtt := 4000.0
	if resBody.Lbtt != expectedLbtt {
		t.Errorf("handler returned unexpected LBTT: got %v want %v", resBody.Lbtt, expectedLbtt)
	}
}

func TestCalculateLBTTHandler_AdditionalDwelling(t *testing.T) {
	reqBody := RequestBody{Price: 500000, IsFirstTimeBuyer: false, IsAdditionalDwelling: true}
	body, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", "/calculate", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(calculateLBTTHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var resBody ResponseBody
	err = json.NewDecoder(rr.Body).Decode(&resBody)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	expectedLbtt := 53350.0
	if resBody.Lbtt != expectedLbtt {
		t.Errorf("handler returned unexpected LBTT: got %v want %v", resBody.Lbtt, expectedLbtt)
	}
}

func TestCalculateLBTTHandler_NegativePrice(t *testing.T) {
	reqBody := RequestBody{Price: -100000, IsFirstTimeBuyer: false, IsAdditionalDwelling: false}
	body, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", "/calculate", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(calculateLBTTHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code for negative price: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestCalculateLBTTHandler_ZeroPrice(t *testing.T) {
	reqBody := RequestBody{Price: 0, IsFirstTimeBuyer: false, IsAdditionalDwelling: false}
	body, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", "/calculate", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(calculateLBTTHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code for zero price: got %v want %v", status, http.StatusBadRequest)
	}
}
