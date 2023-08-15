package test

import (
	conversion2 "conversor_api/internal/conversion"
	"testing"
)

func TestConvert(t *testing.T) {
	ratesProvider := conversion2.NewExchangeRates() // Use conversion.NewExchangeRateProvider()
	ratesProvider.AddRate("BRL", 1.0)
	ratesProvider.AddRate("USD", 4.50)
	converter := conversion2.NewConverter(ratesProvider) // Use conversion.NewConverter()

	// Teste de conversão de BRL para USD
	amount := 10.0
	fromCurrency := "BRL"
	toCurrency := "USD"
	expectedConvertedAmount := 45.0
	convertedAmount, err := converter.Convert(amount, fromCurrency, toCurrency)
	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}
	if convertedAmount != expectedConvertedAmount {
		t.Errorf("Valor convertido incorreto. Esperado: %.2f, Obtido: %.2f", expectedConvertedAmount, convertedAmount)
	}
}
