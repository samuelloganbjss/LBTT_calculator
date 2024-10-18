package calculator

type TaxBandFactory struct{}

func (f TaxBandFactory) CreateCalculator(isFirstTimeBuyer bool, isAdditionalDwelling bool, price float64) *Calculator {
	if isFirstTimeBuyer && isAdditionalDwelling {
		return nil
	}

	var taxBands []TaxBand

	if isFirstTimeBuyer {
		taxBands = []TaxBand{
			BoundedTaxBand{LowerLimit: 175000, UpperLimit: 250000, Rate: 0.02},
			BoundedTaxBand{LowerLimit: 250000, UpperLimit: 325000, Rate: 0.05},
			BoundedTaxBand{LowerLimit: 325000, UpperLimit: 750000, Rate: 0.10},
			UnboundedTaxBand{LowerLimit: 750000, Rate: 0.12},
		}
	} else {
		taxBands = []TaxBand{
			BoundedTaxBand{LowerLimit: 145000, UpperLimit: 250000, Rate: 0.02},
			BoundedTaxBand{LowerLimit: 250000, UpperLimit: 325000, Rate: 0.05},
			BoundedTaxBand{LowerLimit: 325000, UpperLimit: 750000, Rate: 0.10},
			UnboundedTaxBand{LowerLimit: 750000, Rate: 0.12},
		}
	}

	if isAdditionalDwelling {
		adsTax := FixedTaxBand{FixedAmount: price * 0.06}
		taxBands = append(taxBands, adsTax)
	}

	return &Calculator{TaxBands: taxBands}
}
