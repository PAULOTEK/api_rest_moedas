package services

import (
	"conversor_api/internal/conversion/repository"
	"conversor_api/internal/models"
	"fmt"
)

type ConversionService struct {
	Repository repository.ConversionRepository
}

func NewConversionService(repository repository.ConversionRepository) *ConversionService {
	return &ConversionService{Repository: repository}
}
func (s *ConversionService) Convert(amount float64, fromCurrency, toCurrency string) (*models.ConversionModel, error) {
	// Obter a taxa de conversão
	rate, err := s.GetExchangeRate(fromCurrency, toCurrency)
	if err != nil {
		return nil, err
	}

	// Calcular o valor convertido
	convertedValue := amount * rate

	// Criar uma instância de Conversion para representar os dados da conversão
	newConversion := &models.ConversionModel{
		Amount:         amount,
		FromCurrency:   fromCurrency,
		ToCurrency:     toCurrency,
		Rate:           rate,
		ConvertedValue: convertedValue,
	}

	// Salvar a conversão no banco de dados
	err = s.SaveConversion(newConversion)
	if err != nil {
		return nil, err
	}

	return newConversion, nil
}

func (s *ConversionService) GetExchangeRate(fromCurrency, toCurrency string) (float64, error) {
	// Implementar a lógica para obter a taxa de câmbio real de uma fonte confiável
	// Você pode chamar uma API de terceiros ou usar alguma outra fonte confiável
	// Aqui, vamos usar taxas de conversão predefinidas para fins de ilustração
	predefinedRates := map[string]float64{
		"USD": 1.0,
		"EUR": 0.85,
		"BTC": 0.0001,
	}

	rate, exists := predefinedRates[toCurrency]
	if !exists {
		return 0, fmt.Errorf("currency not supported")
	}

	return rate, nil
}

func (s *ConversionService) SaveConversion(conversion *models.ConversionModel) error {
	return s.Repository.Save(*conversion)
}
