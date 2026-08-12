package dto

type CajaDto struct {
	MontoInicial float64 `bson:"montoInicial" json:"montoInicial" validate:"gte=0"`
}

type CerrarCajaDto struct {
	MontoTotal float64 `bson:"montoTotal" json:"montoTotal" validate:"gte=0"`
}
