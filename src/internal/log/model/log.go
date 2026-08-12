package model

import (
	"prestoBackend/src/app/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Log struct {
	ID          bson.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty"`
	Usuario     bson.ObjectID   `bson:"usuario" json:"usuario"`
	Username    string          `bson:"username" json:"username"`
	Accion      enum.AccionEnum `bson:"accion" json:"accion"`
	Descripcion string          `bson:"descripcion" json:"descripcion"`
	Modulo      string          `bson:"modulo" json:"modulo"`
	Fecha       time.Time       `bson:"fecha" json:"fecha"`
	Flag        enum.FlagE      `bson:"flag"`
}
