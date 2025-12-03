package dto

type TarifaDto struct {
	Nombre string     `json:"nombre" validate:"required"`
	Rango  []RangoDto `json:"rango" validate:"required,dive"`
}

type RangoDto struct {
	Rango1 int     `json:"rango1" validate:"required" `
	Rango2 int     `json:"rango2" validate:"required" `
	Costo  float64 `json:"costo" validate:"required" `
	Iva    float64 `json:"iva"`
}
