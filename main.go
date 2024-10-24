package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type RequestBody struct {
	Price float64 `json:"price"`
}

type ResponseBody struct {
	Lbtt float64 `json:"lbtt"`
}

func calculateLBTT(price float64) float64 {
	// Use your actual LBTT calculation logic here.
	// For now, we'll assume a simple fixed calculation.
	return price * 0.05 // For example, 5% of the property price
}

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins for now
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(w) // Enable CORS for each request

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var reqBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil || reqBody.Price <= 0 {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	lbtt := calculateLBTT(reqBody.Price)

	resBody := ResponseBody{Lbtt: lbtt}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resBody)
}

func main() {
	// Set up the HTTP server and routes
	http.HandleFunc("/calculate", calculateHandler)

	// Start the server on port 8080
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
