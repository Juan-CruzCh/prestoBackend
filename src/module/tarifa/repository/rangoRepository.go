package repository

import (
	"context"
	"fmt"
	"prestoBackend/src/core/enum"
	"prestoBackend/src/module/tarifa/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type RangoRepository interface {
	CrearRango(rango *model.Rango, ctx context.Context) (*mongo.InsertOneResult, error)
	ListarRangoPorTarifa(tarifa *bson.ObjectID, ctx context.Context) (*[]model.Rango, error)
	EliminarRango(rango *bson.ObjectID, ctx context.Context) (*mongo.UpdateResult, error)
}

type rangoRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewRangoRepository(db *mongo.Database) RangoRepository {
	return &rangoRepository{
		db:         db,
		collection: db.Collection("Rango"),
	}

}
func (r *rangoRepository) CrearRango(rango *model.Rango, ctx context.Context) (*mongo.InsertOneResult, error) {
	resultado, err := r.collection.InsertOne(ctx, rango)
	if err != nil {
		return nil, err
	}
	return resultado, nil

}

func (r *rangoRepository) ListarRangoPorTarifa(tarifa *bson.ObjectID, ctx context.Context) (*[]model.Rango, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"tarifa": tarifa, "flag": enum.FlagNuevo}, options.Find().SetSort(bson.D{{Key: "rango1", Value: 1}}))
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var rangos []model.Rango
	err = cursor.All(ctx, &rangos)
	if err != nil {
		return nil, err
	}
	return &rangos, nil

}
func (repository *rangoRepository) EliminarRango(rango *bson.ObjectID, ctx context.Context) (*mongo.UpdateResult, error) {
	var flagEliminado bson.D = bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "flag", Value: enum.FlagEliminado},
		}},
	}
	resultado, err := repository.collection.UpdateOne(ctx, bson.M{"flag": enum.FlagNuevo, "_id": rango}, flagEliminado)
	if err != nil {
		return nil, err
	}
	if resultado.MatchedCount == 0 {
		return nil, fmt.Errorf("El rango no existe")
	}
	return resultado, nil
}
func (repository *rangoRepository) EliminarRangoTarifas(tarifa *bson.ObjectID, ctx context.Context) (*mongo.UpdateResult, error) {
	var flagEliminado bson.D = bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "flag", Value: enum.FlagEliminado},
		}},
	}
	resultado, err := repository.collection.UpdateMany(ctx, bson.M{"flag": enum.FlagNuevo, "tarifa": tarifa}, flagEliminado)
	if err != nil {
		return nil, err
	}
	if resultado.MatchedCount == 0 {
		return nil, fmt.Errorf("El rangos no existen")
	}
	return resultado, nil
}
