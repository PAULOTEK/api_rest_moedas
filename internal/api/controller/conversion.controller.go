package controller

import (
	"conversor_api/internal/conversion/services"
	"conversor_api/internal/models"
)

type ConversionController struct {
	Service *services.ConversionService
}

func NewConversionController(service *services.ConversionService) *ConversionController {
	return &ConversionController{Service: service}
}

func (c *ConversionController) PerformConversion(amount float64, fromCurrency, toCurrency string) (*models.ConversionModel, error) {
	// Realizar a conversão de moedas
	result, err := c.Service.Convert(amount, fromCurrency, toCurrency)
	if err != nil {
		return nil, err
	}

	return result, nil
}
