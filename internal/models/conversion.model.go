package models

type ConversionModel struct {
	ID             int64   `json:"id"`
	Amount         float64 `json:"amount"`
	FromCurrency   string  `json:"fromCurrency"`
	ToCurrency     string  `json:"toCurrency"`
	Rate           float64 `json:"rate"`
	ConvertedValue float64 `json:"convertedValue"`
}
