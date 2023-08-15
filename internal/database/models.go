package database

import "gorm.io/gorm"

type ExchangeRate struct {
	gorm.Model
	From       string
	To         string
	Conversion float64
}
