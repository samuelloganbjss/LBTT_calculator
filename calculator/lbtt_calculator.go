package calculator

type LBTTCalculator struct {
	calculator           TaxCalculator
	adsCalculator        TaxCalculator
	isAdditionalDwelling bool
}

func NewLBTTCalculator(calculator TaxCalculator, adsCalculator TaxCalculator, isAdditionalDwelling bool) *LBTTCalculator {
	return &LBTTCalculator{
		calculator:           calculator,
		adsCalculator:        adsCalculator,
		isAdditionalDwelling: isAdditionalDwelling,
	}
}

func (c *LBTTCalculator) Calculate(price float64) float64 {
	baseTax := c.calculator.CalculateTax(price)

	if c.isAdditionalDwelling {
		adsTax := c.adsCalculator.CalculateTax(price)
		return baseTax + adsTax
	}

	return baseTax
}
