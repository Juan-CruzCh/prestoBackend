package repository

import (
	"context"
	"prestoBackend/src/core/enum"
	"prestoBackend/src/core/utils"
	"prestoBackend/src/module/tarifa/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type TarifaRepository interface {
	CrearTarifa(tarifa *model.Tarifa, ctx context.Context) (*mongo.InsertOneResult, error)
	VerificarTarifa(nombre string, ctx context.Context) (int, error)
	ListarTarifasConRagos(ctx context.Context) (*[]bson.M, error)
	ListarTarifas(ctx context.Context) (*[]bson.M, error)
}

type tarifaRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewTarifaRepository(db *mongo.Database) TarifaRepository {
	return &tarifaRepository{
		db:         db,
		collection: db.Collection("Tarifa"),
	}

}
func (r *tarifaRepository) CrearTarifa(tarifa *model.Tarifa, ctx context.Context) (*mongo.InsertOneResult, error) {
	resultado, err := r.collection.InsertOne(ctx, tarifa)
	if err != nil {
		return nil, err
	}
	return resultado, nil
}

func (r *tarifaRepository) VerificarTarifa(nombre string, ctx context.Context) (int, error) {
	cantidad, err := r.collection.CountDocuments(ctx, bson.M{"nombre": nombre})
	if err != nil {
		return 0, err
	}
	return int(cantidad), nil
}

func (r *tarifaRepository) ListarTarifasConRagos(ctx context.Context) (*[]bson.M, error) {
	var pipeline mongo.Pipeline = mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "flag", Value: enum.FlagNuevo},
			}},
		},

		utils.Lookup("Rango", "_id", "tarifa", "rango"),
	}
	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var tarifas []bson.M
	err = cursor.All(ctx, &tarifas)
	if err != nil {
		return nil, err
	}

	return &tarifas, nil
}

func (r *tarifaRepository) ListarTarifas(ctx context.Context) (*[]bson.M, error) {

	cursor, err := r.collection.Find(ctx, bson.M{"flag": enum.FlagNuevo})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var tarifas []bson.M
	err = cursor.All(ctx, &tarifas)
	if err != nil {
		return nil, err
	}

	return &tarifas, nil
}
