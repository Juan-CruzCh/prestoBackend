package repository

import (
	"context"
	"prestoBackend/src/module/pago/model"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type DetallePagoRepository interface {
	CrearDetalle(detalle *model.DetallePago, ctx context.Context) (*mongo.InsertOneResult, error)
}

type detallePagoRepository struct {
	bd         *mongo.Database
	collection *mongo.Collection
}

func NewDetallePagoRepository(bd *mongo.Database) DetallePagoRepository {
	return &detallePagoRepository{
		bd:         bd,
		collection: bd.Collection("DetallePago"),
	}

}

func (repo *detallePagoRepository) CrearDetalle(detalle *model.DetallePago, ctx context.Context) (*mongo.InsertOneResult, error) {

	resultado, err := repo.collection.InsertOne(ctx, detalle)
	if err != nil {
		return nil, err
	}
	return resultado, nil

}
