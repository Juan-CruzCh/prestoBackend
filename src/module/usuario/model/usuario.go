package model

import (
	"prestoBackend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Usuario struct {
	ID              bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Ci              string        `bson:"ci" json:"ci"`
	Nombre          string        `bson:"nombre" json:"nombre"`
	Celular         string        `bson:"celular" json:"celular"`
	ApellidoMaterno string        `bson:"apellidoMaterno" json:"apellidoMaterno"`
	ApellidoPaterno string        `bson:"apellidoPaterno" json:"apellidoPaterno"`
	Usuario         string        `bson:"usuario" json:"usuario"`
	Password        string        `bson:"password" json:"password"`
	Direccion       string        `bson:"direccion" json:"direccion"`
	Flag            enum.FlagE    `bson:"flag" json:"flag"`
	Rol             enum.RolE     `bson:"rol" json:"rol"`
	Fecha           time.Time     `bson:"fecha" json:"fecha"`
}
