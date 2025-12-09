package model

import (
	"prestoBackend/src/core/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type DetallePago struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	Lectura     bson.ObjectID `bson:"lectura"`
	CostoPagado float64       `bson:"costoPagado"`
	Pago        bson.ObjectID `bson:"pago"`
	Flag        enum.FlagE    `bson:"flag"`
	Fecha       time.Time     `bson:"fecha"`
}
