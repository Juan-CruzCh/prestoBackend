package model

import (
	"prestoBackend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Lectura struct {
	ID               bson.ObjectID      `bson:"_id,omitempty"`
	Codigo           string             `bson:"codigo"`
	NumeroLectura    int                `bson:"numeroLectura"`
	Mes              string             `bson:"mes"`
	LecturaActual    int                `bson:"lecturaActual"`
	LecturaAnterior  int                `bson:"lecturaAnterior"`
	ConsumoTotal     int                `bson:"consumoTotal"`
	CostoAPagar      float64            `bson:"costoApagar"`
	Gestion          string             `bson:"gestion"`
	Estado           enum.EstadoLectura `bson:"estado"`
	Medidor          bson.ObjectID      `bson:"medidor"`
	Usuario          bson.ObjectID      `bson:"usuario"`
	Flag             enum.FlagE         `bson:"flag"`
	FechaVencimiento time.Time          `bson:"fechaVencimiento"`
	Fecha            time.Time          `bson:"fecha"`
}
