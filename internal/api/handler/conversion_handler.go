package handler

import (
	"conversor_api/internal/api/controller"
	"conversor_api/internal/api/validation"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ConversionHandler struct {
	Controller *controller.ConversionController
}

func NewConversionHandler(controller *controller.ConversionController) *ConversionHandler {
	return &ConversionHandler{Controller: controller}
}

func (h *ConversionHandler) ConvertHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	amountStr := vars["amount"]
	fromCurrency := vars["from"]
	toCurrency := vars["to"]

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	if !validation.ValidateCurrencyCode(fromCurrency) || !validation.ValidateCurrencyCode(toCurrency) {
		http.Error(w, "Invalid currency code", http.StatusBadRequest)
		return
	}

	if !validation.ValidateAmount(amount) {
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return
	}

	conversionResult, err := h.Controller.PerformConversion(amount, fromCurrency, toCurrency)
	if err != nil {
		http.Error(w, "Conversion error", http.StatusInternalServerError)
		return
	}

	apiResponse := NewConversionResponse(conversionResult.ConvertedValue, toCurrency)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(apiResponse)
	if err != nil {
		http.Error(w, "JSON encoding error", http.StatusInternalServerError)
	}
}
