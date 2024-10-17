package calculator

import "testing"

func TestCalculateLBTT_WithAdditionalDwellingSupplement(t *testing.T) {
	price := 300000.00
	expected := 4600.00 + 18000.00

	calculator := NewLBTTCalculator(StandardTaxCalculator{}, AdditionalDwellingCalculator{}, true)
	result := calculator.Calculate(price)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestCalculateLBTT_WithoutAdditionalDwellingSupplement(t *testing.T) {
	price := 300000.00
	expected := 4600.00

	calculator := NewLBTTCalculator(StandardTaxCalculator{}, AdditionalDwellingCalculator{}, false)
	result := calculator.Calculate(price)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestCalculateLBTT_WithFirstTimeBuyer(t *testing.T) {
	price := 200000.00
	expected := 500.00

	calculator := NewLBTTCalculator(FirstTimeBuyerTaxCalculator{}, AdditionalDwellingCalculator{}, false)
	result := calculator.Calculate(price)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}
