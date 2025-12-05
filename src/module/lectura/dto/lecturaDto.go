package dto

import "go.mongodb.org/mongo-driver/v2/bson"

type LecturaDto struct {
	Mes             string        `json:"mes" validate:"required"`
	LecturaActual   int           `json:"lecturaActual" validate:"required"`
	LecturaAnterior int           `json:"lecturaAnterior" validate:"required"`
	Gestion         string        `json:"gestion" validate:"required"`
	Medidor         bson.ObjectID `json:"medidor" validate:"required"`
}
