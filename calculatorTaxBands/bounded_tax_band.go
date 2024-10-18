package calculator

type BoundedTaxBand struct {
	LowerLimit float64
	UpperLimit float64
	Rate       float64
}

func (b BoundedTaxBand) CalculateTax(price float64) float64 {
	if price <= b.LowerLimit {
		return 0
	}

	upper := price
	if price > b.UpperLimit {
		upper = b.UpperLimit
	}

	taxableAmount := upper - b.LowerLimit
	return taxableAmount * b.Rate
}
