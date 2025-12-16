package model

import (
	"prestoBackend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Lectura struct {
	ID               bson.ObjectID      `bson:"_id,omitempty" json:"_id"`
	Codigo           string             `bson:"codigo" json:"codigo"`
	NumeroLectura    int                `bson:"numeroLectura" json:"numeroLectura"`
	Mes              string             `bson:"mes" json:"mes"`
	LecturaActual    int                `bson:"lecturaActual" json:"lecturaActual"`
	LecturaAnterior  int                `bson:"lecturaAnterior" json:"lecturaAnterior"`
	ConsumoTotal     int                `bson:"consumoTotal" json:"consumoTotal"`
	CostoAPagar      float64            `bson:"costoApagar" json:"costoApagar"`
	Gestion          string             `bson:"gestion" json:"gestion"`
	Estado           enum.EstadoLectura `bson:"estado" json:"estado"`
	Medidor          bson.ObjectID      `bson:"medidor" json:"medidor"`
	Usuario          bson.ObjectID      `bson:"usuario" json:"usuario"`
	Flag             enum.FlagE         `bson:"flag" json:"flag"`
	FechaVencimiento time.Time          `bson:"fechaVencimiento" json:"fechaVencimiento"`
	Fecha            time.Time          `bson:"fecha" json:"fecha"`
}
