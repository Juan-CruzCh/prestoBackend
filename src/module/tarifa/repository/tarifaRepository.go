package repository

import (
	"context"
	"prestoBackend/src/module/tarifa/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type TarifaRepository interface {
	CrearTarifa(tarifa *model.Tarifa, ctx context.Context) (*mongo.InsertOneResult, error)
	VerificarTarifa(nombre string, ctx context.Context) (int, error)
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
