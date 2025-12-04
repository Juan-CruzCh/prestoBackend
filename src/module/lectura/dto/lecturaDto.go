package dto

type LecturaDto struct {
	Mes             string `json:"mes" validate:"required"`
	LecturaActual   int    `json:"lecturaActual" validate:"required"`
	LecturaAnterior int    `json:"lecturaAnterior" validate:"required"`
	Gestion         string `json:"gestion" validate:"required"`
	Medidor         string `json:"medidor" validate:"required"`
}
