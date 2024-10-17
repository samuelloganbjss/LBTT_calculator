package calculator

type FirstTimeBuyerTaxCalculator struct{}

func (t FirstTimeBuyerTaxCalculator) CalculateTax(price float64) float64 {
	var tax float64 = 0
	baseLimit := 175000.00

	if price <= baseLimit {
		return tax
	}

	for _, band := range firstTimeBuyerTaxBands {
		if price > band.LowerLimit {
			upper := price
			if band.UpperLimit != -1 && price > band.UpperLimit {
				upper = band.UpperLimit
			}
			taxableAmount := upper - band.LowerLimit
			tax += taxableAmount * band.Rate
		}
	}

	return tax
}
