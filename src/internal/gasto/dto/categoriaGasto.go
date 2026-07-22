package dto

type CategoriaGastoDto struct {
	Nombre string `json:"nombre" validate:"required"`
}
