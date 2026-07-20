package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type CajaChica struct {
	ID    bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Fecha time.Time     `bson:"fecha" json:"fecha"`
}
