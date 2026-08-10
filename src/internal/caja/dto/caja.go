package dto

type CajaDto struct {
	MontoInicial float64 `bson:"montoInicial" json:"montoInicial" validate:"gte=0"`
}
