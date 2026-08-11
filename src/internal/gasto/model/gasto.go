package model

import (
	"prestoBackend/src/app/enum"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Gasto struct {
	ID             bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Codigo         string        `bson:"codigo" json:"codigo"`
	Descripcion    string        `bson:"descripcion" json:"descripcion"`
	Monto          float64       `bson:"monto" json:"monto"`
	CategoriaGasto bson.ObjectID `bson:"categoriaGasto" json:"categoriaGasto"`
	CajaChica      bson.ObjectID `bson:"cajaChica" json:"cajaChica"`
	Usuario        bson.ObjectID `bson:"usuario" json:"usuario"`
	Comprobante    string        `bson:"comprobante" json:"comprobante"`
	Flag           enum.FlagE    `bson:"flag" json:"flag"`
	Fecha          time.Time     `bson:"fecha" json:"fecha"`
}
