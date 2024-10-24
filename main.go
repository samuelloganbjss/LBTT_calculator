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
	factory := calculator.TaxBandFactory{}
	taxBands := factory.CreateTaxBands(isFirstTimeBuyer, isAdditionalDwelling, price)

	calc := calculator.NewCalculator(taxBands)

	return calc.Calculate(price)
}

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func calculateLBTTHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(w)

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

	lbtt := calculateLBTT(reqBody.Price, reqBody.IsFirstTimeBuyer, reqBody.IsAdditionalDwelling)

	resBody := ResponseBody{Lbtt: lbtt}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resBody)
}

func main() {
	http.HandleFunc("/calculate", calculateLBTTHandler)

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
