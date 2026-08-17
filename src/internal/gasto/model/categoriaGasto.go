package model

import (
	"prestoBackend/src/app/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type CategoriaGasto struct {
	ID     bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nombre string        `bson:"nombre" json:"nombre"`
	Flag   enum.FlagE    `bson:"flag" json:"flag"`
	Fecha  time.Time     `bson:"fecha" json:"fecha"`
}
