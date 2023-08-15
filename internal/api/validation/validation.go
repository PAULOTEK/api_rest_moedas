package validation

import (
	"regexp"
)

func ValidateCurrencyCode(currencyCode string) bool {
	// Validar o código de moeda usando uma expressão regular
	// Neste exemplo, consideraremos códigos de moeda com 3 letras maiúsculas
	match, _ := regexp.MatchString("^[A-Z]{3}$", currencyCode)
	return match
}

func ValidateAmount(amount float64) bool {
	// Validar se o valor é positivo
	return amount >= 0
}
