package repository

import (
	"context"
	"prestoBackend/src/app/common"
	"prestoBackend/src/app/enum"
	"prestoBackend/src/internal/gasto/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type CategoriaGasto interface {
	CrearCategoriaGasto(categoriaGasto *model.CategoriaGasto, ctx context.Context) error
	ListarCategoriaGasto(ctx context.Context) (*[]model.CategoriaGasto, error)
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
	categoriaGasto.Flag = enum.FlagNuevo
	categoriaGasto.Fecha = common.FechaHoraBolivia()
	_, err := r.collection.InsertOne(ctx, categoriaGasto)
	if err != nil {
		return err
	}
	return nil
}

func (r *categoriaGasto) ListarCategoriaGasto(ctx context.Context) (*[]model.CategoriaGasto, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"flag": enum.FlagNuevo})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var restultado []model.CategoriaGasto = []model.CategoriaGasto{}
	for cursor.Next(ctx) {
		var categoria model.CategoriaGasto
		err = cursor.Decode(&categoria)
		if err != nil {
			return nil, err
		}
		restultado = append(restultado, categoria)
	}
	err = cursor.Err()
	if err != nil {
		return nil, err
	}
	return &restultado, nil
}

func (r *categoriaGasto) ActualizarCategoriaGasto(id *bson.ObjectID, categoriaGasto *model.CategoriaGasto, ctx context.Context) error {
	return nil
}

func (r *categoriaGasto) EliminarCategoriaGasto(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
