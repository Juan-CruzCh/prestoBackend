package dto

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type LecturaDto struct {
	Mes             string        `json:"mes" validate:"required"`
	LecturaActual   int           `json:"lecturaActual" validate:"gte=0"`
	LecturaAnterior int           `json:"lecturaAnterior" validate:"gte=0"`
	Gestion         int           `json:"gestion" validate:"required"`
	Medidor         bson.ObjectID `json:"medidor" validate:"required"`
}

type BuscadorLecturaDto struct {
	FechaInicio string `json:"fechaInicio"  validate:"required"`
	FechaFin    string `json:"fechaFin"  validate:"required"`
}
