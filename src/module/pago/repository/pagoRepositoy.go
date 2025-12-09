package repository

import (
	"context"
	"errors"
	"prestoBackend/src/module/pago/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type PagoRepository interface {
	CrearPago(pago *model.Pago, cxt context.Context) (*mongo.InsertOneResult, error)
	CantidadDePagos(cxt context.Context) (int, error)
}

type pagoRepository struct {
	bd         *mongo.Database
	collection *mongo.Collection
}

func NewPagoRepository(bd *mongo.Database) PagoRepository {
	return &pagoRepository{
		bd:         bd,
		collection: bd.Collection("Pago"),
	}

}

func (repo *pagoRepository) CrearPago(pago *model.Pago, cxt context.Context) (*mongo.InsertOneResult, error) {
	resultado, err := repo.collection.InsertOne(cxt, pago)
	if err != nil {
		return nil, errors.New("ocurrio un error al realizar el pag")
	}
	return resultado, nil

}

func (repo *pagoRepository) CantidadDePagos(cxt context.Context) (int, error) {
	cantidad, err := repo.collection.CountDocuments(cxt, bson.M{})
	if err != nil {
		return 0, errors.New("ocurrio un error al realizar el pag")
	}
	cantidad += 1
	return int(cantidad), nil

}
