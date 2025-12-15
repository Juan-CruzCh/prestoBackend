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

type BuscadorMedidorClienteDto struct {
	Pagina          int
	Limite          int
	Nombre          string
	Codigo          string
	ApellidoPaterno string
	ApellidoMaterno string
	Ci              string
	Direccion       string
	NumeroMedidor   string
	Tarifa          string
	Estado          string
	EstadoMedidor   string
}

type MedidorClienteProject struct {
	Nombre          string        `bson:"nombre"`
	ApellidoPaterno string        `bson:"apellidoPaterno"`
	ApellidoMaterno string        `bson:"apellidoMaterno"`
	NumeroMedidor   string        `bson:"numeroMedidor"`
	Estado          string        `bson:"estado"`
	ID              bson.ObjectID `bson:"_id"`
}
