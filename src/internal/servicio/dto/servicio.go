package dto

import "go.mongodb.org/mongo-driver/v2/bson"

type ServicioDto struct {
	Nombre      string  `json:"nombre" validate:"required"`
	Costo       float64 `json:"costo" validate:"required"`
	Descripcion string  `json:"descripcion" validate:"required"`
}

type DetalleServicioDto struct {
	Servicio bson.ObjectID `bson:"servicio" json:"servicio" validate:"required"`
	Tarifa   bson.ObjectID `bson:"tarifa" json:"tarifa" validate:"required"`
}
