package repository

import (
	"context"
	"prestoBackend/src/internal/gasto/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type CategoriaGasto interface {
	CrearCategoriaGasto(categoriaGasto *model.CategoriaGasto, ctx context.Context) error
	ListarCategoriaGasto(ctx context.Context) (interface{}, error)
	ActualizarCategoriaGasto(id *bson.ObjectID, categoriaGasto *model.CategoriaGasto, ctx context.Context) error
	EliminarCategoriaGasto(id *bson.ObjectID, ctx context.Context) error
}

type categoriaGasto struct {
	collection *mongo.Collection
}

func NewCategoriaGastoRepository(db *mongo.Database) *categoriaGasto {
	collection := db.Collection("CategoriaGasto")
	return &categoriaGasto{collection: collection}
}

func (r *categoriaGasto) CrearCategoriaGasto(categoriaGasto *model.CategoriaGasto, ctx context.Context) error {
	return nil
}

func (r *categoriaGasto) ListarCategoriaGasto(ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (r *categoriaGasto) ActualizarCategoriaGasto(id *bson.ObjectID, categoriaGasto *model.CategoriaGasto, ctx context.Context) error {
	return nil
}

func (r *categoriaGasto) EliminarCategoriaGasto(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
