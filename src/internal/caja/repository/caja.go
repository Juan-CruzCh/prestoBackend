package repository

import (
	"context"
	"prestoBackend/src/internal/caja/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Caja interface {
	CrearCaja(caja *model.Caja, ctx context.Context) error
	ListarCaja(ctx context.Context) (interface{}, error)
	ActualizarCaja(id *bson.ObjectID, caja *model.Caja, ctx context.Context) error
	EliminarCaja(id *bson.ObjectID, ctx context.Context) error
}

type caja struct {
	collection *mongo.Collection
}

func NewCajaRepository(db *mongo.Database) *caja {
	collection := db.Collection("Caja")
	return &caja{collection: collection}
}

func (r *caja) CrearCaja(caja *model.Caja, ctx context.Context) error {
	return nil
}

func (r *caja) ListarCaja(ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (r *caja) ActualizarCaja(id *bson.ObjectID, caja *model.Caja, ctx context.Context) error {
	return nil
}

func (r *caja) EliminarCaja(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
