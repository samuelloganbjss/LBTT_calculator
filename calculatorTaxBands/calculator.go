package calculator

type Calculator struct {
	TaxBands []TaxBand
}

func (c Calculator) Calculate(price float64) float64 {
	totalTax := 0.0
	for _, band := range c.TaxBands {
		totalTax += band.CalculateTax(price)
	}
	return totalTax
}
