package dto

type CajaChicaDto struct {
	MontoInicial float64 `json:"montoInicial" binding:"required"`
}
