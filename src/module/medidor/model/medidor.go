package model

import (
	"prestoBackend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Medidor struct {
	ID                 bson.ObjectID      `bson:"_id,omitempty"`
	Codigo             string             `bson:"codigo"`
	NumeroMedidor      string             `bson:"numeroMedidor"`
	Descripcion        string             `bson:"descripcion"`
	Estado             enum.EstadoMedidor `bson:"estado"`
	Cliente            bson.ObjectID      `bson:"cliente"`
	Tarifa             bson.ObjectID      `bson:"tarifa"`
	Direccion          string             `bson:"direccion"`
	FechaInstalacion   time.Time          `bson:"fechaInstalacion"`
	Flag               enum.FlagE         `bson:"flag"`
	Fecha              time.Time          `bson:"fecha"`
	LecturasPendientes int                `bson:"lecturasPendientes"`
}
