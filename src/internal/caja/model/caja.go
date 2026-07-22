package model

import (
	"prestoBackend/src/app/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Caja struct {
	ID           bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Codigo       string        `bson:"codigo" json:"codigo"`
	MontoInicial float64       `bson:"montoInicial" json:"montoInicial"`
	MontoTotal   float64       `bson:"montoTotal" json:"montoTotal"`
	Usuario      bson.ObjectID `bson:"usuario" json:"usuario"`
	MontoPago    float64       `bson:"montoPago" json:"montoPago"`
	FechaInicio  time.Time     `bson:"fechaInicio" json:"fechaInicio"`
	Estado       enum.CajaEnum `bson:"estado" json:"estado"`
	FechaFin     time.Time     `bson:"fechaFin,omitempty" json:"fechaFin,omitempty"`
	Fecha        time.Time     `bson:"fecha" json:"fecha"`
	Flag         enum.FlagE    `bson:"flag" json:"flag"`
}
