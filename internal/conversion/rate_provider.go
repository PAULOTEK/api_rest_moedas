package conversion

type RateProvider interface {
	AddRate(currency string, rate float64)
	GetRate(currency string) (float64, bool)
}
