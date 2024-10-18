package calculator

type FixedTaxBand struct {
	FixedAmount float64
}

func (f FixedTaxBand) CalculateTax(price float64) float64 {
	return f.FixedAmount
}
