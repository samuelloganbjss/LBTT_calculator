package calculator

type TaxBand interface {
	CalculateTax(price float64) float64
}
