package model

import (
	"prestoBackend/src/app/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type CajaChica struct {
	Codigo        string        `bson:"codigo" json:"codigo"`
	Estado        string        `bson:"estado" json:"estado"`
	MontoInicial  float64       `bson:"montoInicial" json:"montoInicial"`
	MontoActual   float64       `bson:"montoActual" json:"montoActual"`
	Usuario       bson.ObjectID `bson:"usuario" json:"usuario"`
	FechaApertura time.Time     `bson:"fecha_apertura" json:"fecha_apertura"`
	FechaFin      time.Time     `bson:"fechaFin,omitempty" json:"fechaFin,omitempty"`
	FechaInicio   time.Time     `bson:"fechaInicio" json:"fechaInicio"`
	ID            bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Fecha         time.Time     `bson:"fecha" json:"fecha"`
	Flag          enum.FlagE    `bson:"flag"`
}
