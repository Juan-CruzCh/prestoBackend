package model

import (
	"prestoBackend/src/app/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Pago struct {
	ID             bson.ObjectID `bson:"_id,omitempty"`
	NumeroPago     int           `bson:"numeroPago"`
	Total          float64       `bson:"total"`
	Usuario        bson.ObjectID `bson:"usuario"`
	TipoPago       enum.TipoPago `bson:"tipoPago"`
	Flag           enum.FlagE    `bson:"flag"`
	Fecha          time.Time     `bson:"fecha"`
	FechaAnulacion time.Time     `bson:"fechaAnulacion"`
	Cliente        bson.ObjectID `bson:"cliente"`
	Medidor        bson.ObjectID `bson:"medidor"`
	Caja           bson.ObjectID `bson:"caja"`
}
