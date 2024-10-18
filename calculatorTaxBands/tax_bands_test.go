package calculator

import "testing"

func TestBoundedTaxBand(t *testing.T) {
	band := BoundedTaxBand{LowerLimit: 145000, UpperLimit: 250000, Rate: 0.02}

	price := 200000.00
	expected := (200000 - 145000) * 0.02
	result := band.CalculateTax(price)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestUnboundedTaxBand(t *testing.T) {
	band := UnboundedTaxBand{LowerLimit: 750000, Rate: 0.12}

	price := 1000000.00
	expected := (1000000 - 750000) * 0.12
	result := band.CalculateTax(price)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestFixedTaxBand(t *testing.T) {
	band := FixedTaxBand{FixedAmount: 500.00}

	price := 1000000.00
	expected := 500.00
	result := band.CalculateTax(price)

	if result != expected {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestFactoryHandlesFirstTimeBuyerAndAdditionalDwellingConflict(t *testing.T) {
	factory := TaxBandFactory{}
	_, err := factory.CreateTaxBands(true, true)

	if err == nil {
		t.Errorf("Expected an error for first-time buyer and additional dwelling, but got nil")
	}
}

func TestFactoryCreatesFirstTimeBuyerTaxBands(t *testing.T) {
	factory := TaxBandFactory{}
	taxBands, err := factory.CreateTaxBands(true, false)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(taxBands) != 4 {
		t.Errorf("Expected 4 tax bands for first-time buyer, got %d", len(taxBands))
	}
}

func TestFactoryCreatesStandardWithADSTaxBands(t *testing.T) {
	factory := TaxBandFactory{}
	taxBands, err := factory.CreateTaxBands(false, true)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(taxBands) != 5 {
		t.Errorf("Expected 5 tax bands for standard buyer with ADS, got %d", len(taxBands))
	}
}

func TestFullTaxCalculationForStandardBuyer(t *testing.T) {
	factory := TaxBandFactory{}
	price := 300000.00
	isFirstTimeBuyer := false
	isAdditionalDwelling := false

	taxBands, err := factory.CreateTaxBands(isFirstTimeBuyer, isAdditionalDwelling)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	totalTax := 0.0
	for _, band := range taxBands {
		totalTax += band.CalculateTax(price)
	}

	expectedTax := 4600.00
	if totalTax != expectedTax {
		t.Errorf("Expected total tax %f but got %f", expectedTax, totalTax)
	}
}

func TestFullTaxCalculationForFirstTimeBuyer(t *testing.T) {
	factory := TaxBandFactory{}
	price := 300000.00
	isFirstTimeBuyer := true
	isAdditionalDwelling := false

	taxBands, err := factory.CreateTaxBands(isFirstTimeBuyer, isAdditionalDwelling)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	totalTax := 0.0
	for _, band := range taxBands {
		totalTax += band.CalculateTax(price)
	}

	expectedTax := 4000.00
	if totalTax != expectedTax {
		t.Errorf("Expected total tax %f but got %f", expectedTax, totalTax)
	}
}
