package calculator

type TaxBand struct {
	LowerLimit float64
	UpperLimit float64
	Rate       float64
}

var taxBands = []TaxBand{
	{LowerLimit: 145000, UpperLimit: 250000, Rate: 0.02},
	{LowerLimit: 250000, UpperLimit: 325000, Rate: 0.05},
	{LowerLimit: 325000, UpperLimit: 750000, Rate: 0.10},
	{LowerLimit: 750000, UpperLimit: -1, Rate: 0.12},
}

var firstTimeBuyerTaxBands = []TaxBand{
	{LowerLimit: 175000, UpperLimit: 250000, Rate: 0.02},
	{LowerLimit: 250000, UpperLimit: 325000, Rate: 0.05},
	{LowerLimit: 325000, UpperLimit: 750000, Rate: 0.10},
	{LowerLimit: 750000, UpperLimit: -1, Rate: 0.12},
}
