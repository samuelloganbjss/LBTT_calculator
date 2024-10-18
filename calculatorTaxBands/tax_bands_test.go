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
	err := factory.CreateTaxBands(true, true, 500000)

	if err != nil {
		t.Errorf("Expected an error for first-time buyer and additional dwelling, but got nil")
	}
}

func TestFactoryCreatesFirstTimeBuyerTaxBands(t *testing.T) {
	factory := TaxBandFactory{}
	taxBands := factory.CreateTaxBands(true, false, 500000)

	if len(taxBands) != 4 {
		t.Errorf("Expected 4 tax bands for first-time buyer, got %d", len(taxBands))
	}
}

func TestFactoryCreatesStandardWithADSTaxBands(t *testing.T) {
	factory := TaxBandFactory{}
	taxBands := factory.CreateTaxBands(false, true, 500000)

	if len(taxBands) != 5 {
		t.Errorf("Expected 5 tax bands for standard buyer with ADS, got %d", len(taxBands))
	}
}

func TestFullTaxCalculationForStandardBuyer(t *testing.T) {
	factory := TaxBandFactory{}
	price := 300000.00
	isFirstTimeBuyer := false
	isAdditionalDwelling := false
	taxBands := factory.CreateTaxBands(isFirstTimeBuyer, isAdditionalDwelling, price)
	calculator := NewCalculator(taxBands)

	totalTax := calculator.Calculate(price)

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
	taxBands := factory.CreateTaxBands(isFirstTimeBuyer, isAdditionalDwelling, price)
	calculator := NewCalculator(taxBands)

	totalTax := calculator.Calculate(price)

	expectedTax := 4000.00
	if totalTax != expectedTax {
		t.Errorf("Expected total tax %f but got %f", expectedTax, totalTax)
	}
}

func TestFullTaxCalculationForAdditionalDwellingSupplement(t *testing.T) {
	factory := TaxBandFactory{}
	price := 500000.00
	isFirstTimeBuyer := false
	isAdditionalDwelling := true
	taxBands := factory.CreateTaxBands(isFirstTimeBuyer, isAdditionalDwelling, price)
	calculator := NewCalculator(taxBands)

	totalTax := calculator.Calculate(price)

	expectedTax := 53350.0
	if totalTax != expectedTax {
		t.Errorf("Expected total tax %f but got %f", expectedTax, totalTax)
	}
}
