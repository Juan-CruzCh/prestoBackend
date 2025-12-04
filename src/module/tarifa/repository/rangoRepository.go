package repository

import (
	"context"
	"prestoBackend/src/module/tarifa/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type RangoRepository interface {
	CrearRango(rango *model.Rango, ctx context.Context) (*mongo.InsertOneResult, error)
	ListarRangoPorTarifa(tarifa *bson.ObjectID, ctx context.Context) (*[]model.Rango, error)
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
	cursor, err := r.collection.Find(ctx, bson.M{"tarifa": tarifa})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var rangos []model.Rango
	err = cursor.All(ctx, &rangos)

	return &rangos, nil

}
