package repository

import (
	"context"
	"prestoBackend/src/app/database/aggregation"
	"prestoBackend/src/app/enum"
	"prestoBackend/src/internal/gasto/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Gasto interface {
	CrearGasto(gasto *model.Gasto, ctx context.Context) error
	ListarGasto(ctx context.Context) (*[]bson.M, error)
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
	_, err := r.collection.InsertOne(ctx, gasto)
	if err != nil {
		return err
	}
	return nil
}

func (r *gasto) ListarGasto(ctx context.Context) (*[]bson.M, error) {
	var pipeline mongo.Pipeline = mongo.Pipeline{
		bson.D{
			{
				Key: "$match", Value: bson.D{
					{
						Key: "flag", Value: enum.FlagNuevo,
					},
				},
			},
		},
		aggregation.Lookup("CategoriaGasto", "categoriaGasto", "_id", "categoriaGasto"),
		aggregation.Lookup("CajaChica", "cajaChica", "_id", "cajaChica"),
		aggregation.Lookup("Usuario", "usuario", "_id", "usuario"),

		bson.D{
			{
				Key: "$project", Value: bson.D{
					{
						Key: "_id", Value: 1,
					},
					{
						Key: "codigo", Value: 1,
					},

					{
						Key: "descripcion", Value: 1,
					},
					{
						Key: "monto", Value: 1,
					},
					{
						Key: "comprobante", Value: 1,
					},
					{
						Key: "fecha", Value: 1,
					},
					{
						Key: "usuario", Value: aggregation.ArrayElemAt("$usuario.usuario", 0),
					},
					{
						Key: "categoriaGasto ", Value: aggregation.ArrayElemAt("$categoriaGasto.nombre", 0),
					},
					{
						Key: "cajaChica", Value: aggregation.ArrayElemAt("$cajaChica.codigo", 0),
					},
				},
			},
		},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var gastos []bson.M = []bson.M{}
	for cursor.Next(ctx) {
		var gasto bson.M = bson.M{}
		err = cursor.Decode(&gasto)
		if err != nil {
			return nil, err
		}
		gastos = append(gastos, gasto)
	}
	err = cursor.Err()
	if err != nil {
		return nil, err
	}
	return &gastos, nil
}

func (r *gasto) ActualizarGasto(id *bson.ObjectID, gasto *model.Gasto, ctx context.Context) error {
	return nil
}

func (r *gasto) EliminarGasto(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
