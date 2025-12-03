package model

import (
	"prestoBackend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Rango struct {
	ID     bson.ObjectID `bson:"_id,omitempty"`
	Rango1 int           `bson:"rango1"`
	Rango2 int           `bson:"rango2"`
	Costo  float64       `bson:"costo"`
	Tarifa bson.ObjectID `bson:"tarifa"`
	Iva    float64       `bson:"iva"`
	Flag   enum.FlagE    `bson:"flag"`
	Fecha  time.Time     `bson:"fecha"`
}
