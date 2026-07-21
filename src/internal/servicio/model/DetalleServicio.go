package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type DetalleServicio struct {
	ID       bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Servicio bson.ObjectID `bson:"servicio" json:"servicio"`
	Tarifa   bson.ObjectID `bson:"tarifa" json:"tarifa"`
	Fecha    time.Time     `bson:"fecha" json:"fecha"`
}
