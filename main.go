package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TaxRequest struct {
	Price                float64 `json:"price"`
	IsFirstTimeBuyer     bool    `json:"isFirstTimeBuyer"`
	IsAdditionalDwelling bool    `json:"isAdditionalDwelling"`
}

type TaxResponse struct {
	Lbtt float64 `json:"lbtt"`
}

func calculateLBTT(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Handle CORS for frontend requests

	var req TaxRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Example calculation logic (replace with your own LBTT logic)
	lbtt := req.Price * 0.05

	response := TaxResponse{Lbtt: lbtt}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/calculate", calculateLBTT)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
