package repository

import (
	"context"
	"errors"
	"prestoBackend/src/core/enum"
	"prestoBackend/src/core/utils"
	"prestoBackend/src/module/pago/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type PagoRepository interface {
	CrearPago(pago *model.Pago, cxt context.Context) (*mongo.InsertOneResult, error)
	CantidadDePagos(cxt context.Context) (int, error)
	DetallePago(idPago *bson.ObjectID, ctx context.Context) (*bson.M, error)
	BuscarPagoId(idPago *bson.ObjectID, cxt context.Context) (model.Pago, error)
	ListarPagos(ctx context.Context) (*[]bson.M, error)
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

func (repo *pagoRepository) BuscarPagoId(idPago *bson.ObjectID, cxt context.Context) (model.Pago, error) {
	var data model.Pago
	err := repo.collection.FindOne(cxt, bson.M{"_id": idPago, "flag": enum.FlagNuevo}).Decode(&data)
	if err != nil {
		return model.Pago{}, err
	}
	return data, nil

}

func (repo *pagoRepository) DetallePago(idPago *bson.ObjectID, ctx context.Context) (*bson.M, error) {
	var pipepine mongo.Pipeline = mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{
				{
					Key: "_id", Value: idPago,
				},
				{
					Key: "flag", Value: enum.FlagNuevo,
				},
			}},
		},

		utils.Lookup("Cliente", "cliente", "_id", "cliente"),
		utils.Lookup("Medidor", "medidor", "_id", "medidor"),
		utils.Lookup("DetallePago", "_id", "pago", "detallePago"),

		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "numeroPago", Value: 1},
				{Key: "total", Value: 1},
				{Key: "fecha", Value: 1},
				{Key: "numeroMedidor", Value: utils.ArrayElemAt("$medidor.numeroMedidor", 0)},
				{Key: "nombre", Value: utils.ArrayElemAt("$cliente.nombre", 0)},
				{Key: "apellidoPaterno", Value: utils.ArrayElemAt("$cliente.apellidoPaterno", 0)},
				{Key: "apellidoMaterno", Value: utils.ArrayElemAt("$cliente.apellidoMaterno", 0)},
				{Key: "detallePago", Value: 1},
				{Key: "direccion", Value: utils.ArrayElemAt("$medidor.direccion", 0)},
				{Key: "codigoCliente", Value: utils.ArrayElemAt("$cliente.codigo", 0)},
			}},
		},
	}

	cursor, err := repo.collection.Aggregate(ctx, pipepine)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var data []bson.M = []bson.M{}
	err = cursor.All(ctx, &data)
	if err != nil {
		return nil, err
	}
	return &data[0], nil

}

func (repo *pagoRepository) ListarPagos(ctx context.Context) (*[]bson.M, error) {
	var pipepine mongo.Pipeline = mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{

				{
					Key: "flag", Value: enum.FlagNuevo,
				},
			}},
		},
		utils.Lookup("Cliente", "cliente", "_id", "cliente"),
		utils.Lookup("Medidor", "medidor", "_id", "medidor"),
		utils.Lookup("DetallePago", "_id", "pago", "detallePago"),
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "numeroPago", Value: 1},
				{Key: "total", Value: 1},
				{Key: "fecha", Value: 1},
				{Key: "numeroMedidor", Value: utils.ArrayElemAt("$medidor.numeroMedidor", 0)},
				{Key: "nombre", Value: utils.ArrayElemAt("$cliente.nombre", 0)},
				{Key: "apellidoPaterno", Value: utils.ArrayElemAt("$cliente.apellidoPaterno", 0)},
				{Key: "apellidoMaterno", Value: utils.ArrayElemAt("$cliente.apellidoMaterno", 0)},
				{Key: "detallePago", Value: 1},
				{Key: "codigoCliente", Value: utils.ArrayElemAt("$cliente.codigo", 0)},
			}},
		},
	}
	cursor, err := repo.collection.Aggregate(ctx, pipepine)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var data []bson.M = []bson.M{}
	err = cursor.All(ctx, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
