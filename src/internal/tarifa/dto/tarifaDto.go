package dto

type TarifaDto struct {
	Nombre string     `json:"nombre" validate:"required"`
	Rango  []RangoDto `json:"rango" validate:"required,dive"`
}

type RangoDto struct {
	Rango1 int     `json:"rango1" validate:"gte=0"`
	Rango2 int     `json:"rango2" validate:"gte=0"`
	Costo  float64 `json:"costo" validate:"gte=0"`
	Iva    float64 `json:"iva" validate:"gte=0"`
}
