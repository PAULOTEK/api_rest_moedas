package repository

import (
	"conversor_api/internal/models"
	"database/sql"
	"fmt"
)

type MySQLRepository struct {
	DB *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{DB: db}
}

func (r *MySQLRepository) Save(conversion models.ConversionModel) error {
	query := "INSERT INTO conversions (amount, from_currency, to_currency, rate, converted_value) VALUES (?, ?, ?, ?, ?)"
	_, err := r.DB.Exec(query, conversion.Amount, conversion.FromCurrency, conversion.ToCurrency, conversion.Rate, conversion.ConvertedValue)
	if err != nil {
		return fmt.Errorf("error saving conversion: %v", err)
	}
	return nil
}

type ConversionRepository interface {
	Save(conversion models.ConversionModel) error
}

func (r *MySQLRepository) GetByID(id int) (*models.ConversionModel, error) {
	// Implementar lógica para obter conversão por ID do banco de dados
	return nil, nil
}

func (r *MySQLRepository) GetAll() ([]models.ConversionModel, error) {
	// Implementar lógica para obter todas as conversões do banco de dados
	return nil, nil
}

func (r *MySQLRepository) GetByCurrencyPair(fromCurrency, toCurrency string) ([]models.ConversionModel, error) {
	// Implementar lógica para obter conversões por par de moedas do banco de dados
	return nil, nil
}
