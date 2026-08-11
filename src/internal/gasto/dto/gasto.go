package dto

import "go.mongodb.org/mongo-driver/v2/bson"

type GastoDto struct {
	Nombre         string        `json:"nombre" validate:"required"`
	Descripcion    string        `json:"descripcion" validate:"required"`
	Monto          float64       `json:"monto" validate:"required"`
	CategoriaGasto bson.ObjectID `json:"categoriaGasto" validate:"required"`
}
