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

func TestFactoryCreatesFirstTimeBuyerTaxBands(t *testing.T) {
	factory := TaxBandFactory{}
	taxBands := factory.CreateTaxBands(true, false)

	if len(taxBands) != 4 {
		t.Errorf("Expected 4 tax bands for first-time buyer, got %d", len(taxBands))
	}
}

func TestFactoryCreatesStandardWithADSTaxBands(t *testing.T) {
	factory := TaxBandFactory{}
	taxBands := factory.CreateTaxBands(false, true)

	if len(taxBands) != 5 {
		t.Errorf("Expected 5 tax bands for standard buyer with ADS, got %d", len(taxBands))
	}
}

func TestFactoryCreatesStandardTaxBands(t *testing.T) {
	factory := TaxBandFactory{}
	taxBands := factory.CreateTaxBands(false, false)

	if len(taxBands) != 4 {
		t.Errorf("Expected 4 tax bands for standard buyer, got %d", len(taxBands))
	}
}
