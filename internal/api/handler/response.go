package handler

type ConversionResponse struct {
	ValorConvertido float64 `json:"valorConvertido"`
	SimboloMoeda    string  `json:"simboloMoeda"`
}

func NewConversionResponse(valorConvertido float64, simboloMoeda string) ConversionResponse {
	return ConversionResponse{
		ValorConvertido: valorConvertido,
		SimboloMoeda:    simboloMoeda,
	}
}
