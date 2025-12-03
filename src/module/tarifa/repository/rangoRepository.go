package repository

import (
	"context"
	"prestoBackend/src/module/tarifa/model"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type RangoRepository interface {
	CrearRango(rango *model.Rango, ctx context.Context) (*mongo.InsertOneResult, error)
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
