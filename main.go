package conversor_api

import (
	"conversor_api/internal/api/controller"
	"conversor_api/internal/api/handler"
	"conversor_api/internal/conversion/repository"
	"conversor_api/internal/conversion/services"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func Main() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/currency_converter")
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	router := mux.NewRouter()

	var conversionRepository repository.ConversionRepository // Não especifique o tipo aqui
	conversionRepository = repository.NewMySQLRepository(db) // Atribua a instância diretamente

	conversionService := services.NewConversionService(conversionRepository)
	conversionController := controller.NewConversionController(conversionService)
	conversionHandler := handler.NewConversionHandler(conversionController)

	router.HandleFunc("/exchange/{amount}/{from}/{to}", conversionHandler.ConvertHandler).Methods("GET")

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
