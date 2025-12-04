package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type MedidorDto struct {
	NumeroMedidor    string        `json:"numeroMedidor"  validate:"required"`
	Descripcion      string        `json:"descripcion" `
	Cliente          bson.ObjectID `json:"cliente"  validate:"required"`
	Tarifa           bson.ObjectID `json:"tarifa"  validate:"required"`
	Direccion        string        `json:"direccion"  validate:"required"`
	FechaInstalacion time.Time     `json:"fechaInstalacion"  validate:"required"`
}
