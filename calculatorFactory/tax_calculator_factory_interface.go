package calculator

type TaxCalculatorFactory interface {
	CreateTaxCalculator(isFirstTimeBuyer bool, isAdditionalDwelling bool) TaxCalculator
}
