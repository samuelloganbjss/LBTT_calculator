package calculator

import "errors"

type TaxBandFactory struct{}

func (f TaxBandFactory) CreateTaxBands(isFirstTimeBuyer bool, isAdditionalDwelling bool, price float64) ([]TaxBand, error) {
	if isFirstTimeBuyer && isAdditionalDwelling {
		return nil, errors.New("a first-time buyer cannot have an additional dwelling tax")
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

	return taxBands, nil
}
