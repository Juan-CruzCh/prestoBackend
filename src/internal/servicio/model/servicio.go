package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Servicio struct {
	Nombre      string        `bson:"nombre" json:"nombre"`
	Costo       float64       `bson:"costo" json:"costo"`
	Descripcion string        `bson:"descripcion" json:"descripcion"`
	ID          bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Fecha       time.Time     `bson:"fecha" json:"fecha"`
}
