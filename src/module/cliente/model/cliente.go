package model

import (
	"prestoBackend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Cliente struct {
	ID              bson.ObjectID `bson:"_id,omitempty"`
	Codigo          string        `bson:"codigo"`
	Ci              string        `bson:"ci"`
	Nombre          string        `bson:"nombre"`
	ApellidoMaterno string        `bson:"apellidoMaterno"`
	ApellidoPaterno string        `bson:"apellidoPaterno"`
	Flag            enum.FlagE    `bson:"flag"`
	Fecha           time.Time     `bson:"fecha"`
}
