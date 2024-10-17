package calculator

type AdditionalDwellingCalculator struct{}

func (t AdditionalDwellingCalculator) CalculateTax(price float64) float64 {
	return price * 0.06
}
