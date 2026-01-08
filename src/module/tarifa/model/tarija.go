package model

import (
	"prestoBackend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Tarifa struct {
	ID     bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	Nombre string        `bson:"nombre" json:"nombre"`
	Flag   enum.FlagE    `bson:"flag" json:"flag"`
	Fecha  time.Time     `bson:"fecha" json:"fecha"`
}
