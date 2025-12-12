package model

import (
	"prestoBackend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Cliente struct {
	ID              bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	Codigo          string        `bson:"codigo" json:"codigo"`
	Ci              string        `bson:"ci" json:"ci"`
	Nombre          string        `bson:"nombre" json:"nombre"`
	ApellidoMaterno string        `bson:"apellidoMaterno" json:"apellidoMaterno"`
	ApellidoPaterno string        `bson:"apellidoPaterno" json:"apellidoPaterno"`
	Flag            enum.FlagE    `bson:"flag" json:"flag"`
	Fecha           time.Time     `bson:"fecha"  json:"fecha"`
}
