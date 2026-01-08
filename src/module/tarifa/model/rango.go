package model

import (
	"prestoBackend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Rango struct {
	ID     bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Rango1 int           `bson:"rango1" json:"rango1"`
	Rango2 int           `bson:"rango2" json:"rango2"`
	Costo  float64       `bson:"costo" json:"costo"`
	Tarifa bson.ObjectID `bson:"tarifa" json:"tarifa"`
	Iva    float64       `bson:"iva" json:"iva"`
	Flag   enum.FlagE    `bson:"flag" json:"flag"`
	Fecha  time.Time     `bson:"fecha" json:"fecha"`
}
