package main

import (
	calculator "LBTT_Calculator/calculatorTaxBands"
	"encoding/json"
	"log"
	"net/http"
)

type RequestBody struct {
	Price                float64 `json:"price"`
	IsFirstTimeBuyer     bool    `json:"isFirstTimeBuyer"`
	IsAdditionalDwelling bool    `json:"isAdditionalDwelling"`
}

type ResponseBody struct {
	Lbtt float64 `json:"lbtt"`
}

func calculateLBTT(price float64, isFirstTimeBuyer, isAdditionalDwelling bool) float64 {
	// Use the TaxBandFactory to create the appropriate tax bands
	factory := calculator.TaxBandFactory{}
	taxBands := factory.CreateTaxBands(isFirstTimeBuyer, isAdditionalDwelling, price)

	// Create a new calculator with the generated tax bands
	calc := calculator.NewCalculator(taxBands)

	// Calculate the total tax
	return calc.Calculate(price)
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

	// Calculate the LBTT using the calculateLBTT function
	lbtt := calculateLBTT(reqBody.Price, reqBody.IsFirstTimeBuyer, reqBody.IsAdditionalDwelling)

	// Return the LBTT as a JSON response
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
