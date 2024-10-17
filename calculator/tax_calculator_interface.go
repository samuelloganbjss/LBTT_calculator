package calculator

type TaxCalculator interface {
	CalculateTax(price float64) float64
}
