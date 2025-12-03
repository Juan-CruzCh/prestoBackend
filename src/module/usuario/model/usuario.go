package model

import (
	"prestoBackend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Usuario struct {
	ID              bson.ObjectID `bson:"_id,omitempty"`
	Ci              string        `bson:"ci"`
	Nombre          string        `bson:"nombre"`
	Celular         string        `bson:"celular"`
	ApellidoMaterno string        `bson:"apellidoMaterno"`
	ApellidoPaterno string        `bson:"apellidoPaterno"`
	Usuario         string        `bson:"usuario"`
	Password        string        `bson:"password"`
	Direccion       string        `bson:"direccion"`
	Flag            enum.FlagE    `bson:"flag"`
	Rol             enum.RolE     `bson:"rol"`
	Fecha           time.Time     `bson:"fecha"`
}
