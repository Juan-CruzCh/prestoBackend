package dto

type CajaChicaDto struct {
	MontoInicial float64 `json:"montoInicial" binding:"required"`
	FechaInicio  string  `json:"fechaInicio" binding:"required"`
	FechaFin     string  `json:"fechaFin" binding:"required"`
}
