package app

type Input interface {
	GetHousePrice() float64
	IsFirstTimeBuyer() bool
	IsAdditionalDwelling() bool
}
