package calculator

type UnboundedTaxBand struct {
	LowerLimit float64
	Rate       float64
}

func (u UnboundedTaxBand) CalculateTax(price float64) float64 {
	if price <= u.LowerLimit {
		return 0
	}

	taxableAmount := price - u.LowerLimit
	return taxableAmount * u.Rate
}
