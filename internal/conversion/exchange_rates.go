package conversion

type ExchangeRates struct {
	rates map[string]float64
}

func NewExchangeRates() *ExchangeRates {
	return &ExchangeRates{
		make(map[string]float64),
	}
}

func (e *ExchangeRates) AddRate(currency string, rate float64) {
	e.rates[currency] = rate
}

func (e *ExchangeRates) GetRate(currency string) (float64, bool) {
	rate, exists := e.rates[currency]
	return rate, exists
}
