package repository

import (
	"context"
	"prestoBackend/src/internal/gasto/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Gasto interface {
	CrearGasto(gasto *model.Gasto, ctx context.Context) error
	ListarGasto(ctx context.Context) (interface{}, error)
	ActualizarGasto(id *bson.ObjectID, gasto *model.Gasto, ctx context.Context) error
	EliminarGasto(id *bson.ObjectID, ctx context.Context) error
}

type gasto struct {
	collection *mongo.Collection
}

func NewGastoRepository(db *mongo.Database) *gasto {
	collection := db.Collection("Gasto")
	return &gasto{collection: collection}
}

func (r *gasto) CrearGasto(gasto *model.Gasto, ctx context.Context) error {
	return nil
}

func (r *gasto) ListarGasto(ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (r *gasto) ActualizarGasto(id *bson.ObjectID, gasto *model.Gasto, ctx context.Context) error {
	return nil
}

func (r *gasto) EliminarGasto(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
