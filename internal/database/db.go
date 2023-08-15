package database

import (
	"conversor_api/internal/conversion"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase() error {
	dsn := "user:password@tcp(127.0.0.1:3306)/conversor_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Migrate the schema
	db.AutoMigrate(&conversion.ExchangeRates{}) // Use a sintaxe completa para acessar o tipo

	return nil
}

func SaveExchangeRate(rate conversion.ExchangeRates) error {
	result := db.Create(&rate)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
