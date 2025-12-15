package dto

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type LecturaDto struct {
	Mes             string        `json:"mes" validate:"required"`
	LecturaActual   int           `json:"lecturaActual" validate:"gte=0"`
	LecturaAnterior int           `json:"lecturaAnterior" validate:"gte=0"`
	Gestion         string        `json:"gestion" validate:"required"`
	Medidor         bson.ObjectID `json:"medidor" validate:"required"`
}

type BuscadorLecturaDto struct {
	Codigo      string `json:"codigo"`
	FechaInicio string `json:"fechaInicio"  validate:"required"`
	FechaFin    string `json:"fechaFin"  validate:"required"`
}
