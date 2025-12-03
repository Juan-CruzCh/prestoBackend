package model

import (
	"prestoBackend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Tarifa struct {
	ID     bson.ObjectID `bson:"_id,omitempty"`
	Nombre string        `bson:"nombre"`
	Flag   enum.FlagE    `bson:"flag"`
	Fecha  time.Time     `bson:"fecha"`
}
