package dto

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type PagoDto struct {
	Cliente  bson.ObjectID `json:"cliente" validate:"required"`
	Medidor  bson.ObjectID `json:"medidor"  validate:"required"`
	Lecturas []LecturasDto `json:"lecturas" validate:"required,dive,required"`
}

type LecturasDto struct {
	Lectura bson.ObjectID `json:"lectura" validate:"required"`
}

type BuscardorPagoDto struct {
	CodigoCliente   string
	Ci              string
	Nombre          string
	ApellidoMaterno string
	ApellidoPaterno string
	FechaInicio     string
	FechaFin        string
	NumeroMedidor   string
	Pagina          int
	Limite          int
}
