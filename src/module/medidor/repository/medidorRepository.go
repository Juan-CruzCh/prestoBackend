package repository

import (
	"context"
	"errors"
	"fmt"
	"prestoBackend/src/core/enum"
	"prestoBackend/src/module/medidor/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MedidorRepository interface {
	ListarMedidor()
	EliminarMedidor()
	ActualizarMedidor()
	CrearMedidor(medidor *model.Medidor, ctx context.Context) (*mongo.InsertOneResult, error)
	CantidadMedidor(ctx context.Context) (int, error)
	ObtenerMedidor(medidor *bson.ObjectID, ctx context.Context) (*model.Medidor, error)
	ActualizaLecturasPendientesMedidor(cantidad int, medidor *bson.ObjectID, ctx context.Context) error
}

type medidorRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewMedidorRespository(db *mongo.Database) MedidorRepository {
	return &medidorRepository{
		db:         db,
		collection: db.Collection("Medidor"),
	}
}

func (r *medidorRepository) ListarMedidor() {

}

func (r *medidorRepository) EliminarMedidor() {

}

func (r *medidorRepository) ActualizarMedidor() {

}

func (r *medidorRepository) CrearMedidor(medidor *model.Medidor, ctx context.Context) (*mongo.InsertOneResult, error) {

	cantidad, err := r.collection.CountDocuments(ctx, bson.M{"numeroMedidor": medidor.NumeroMedidor})
	if err != nil {
		return nil, err
	}
	if cantidad > 0 {
		return nil, errors.New("el medidor ya se encuentra registrado")
	}
	resultado, err := r.collection.InsertOne(ctx, medidor)
	if err != nil {
		return nil, err
	}
	return resultado, nil

}

func (r *medidorRepository) CantidadMedidor(ctx context.Context) (int, error) {

	cantidad, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}

	return int(cantidad), nil

}

func (r *medidorRepository) ObtenerMedidor(medidor *bson.ObjectID, ctx context.Context) (*model.Medidor, error) {
	var data model.Medidor
	err := r.collection.FindOne(ctx, bson.M{"_id": medidor, "flag": enum.FlagNuevo}).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &data, nil

}

func (r *medidorRepository) ActualizaLecturasPendientesMedidor(cantidad int, medidor *bson.ObjectID, ctx context.Context) error {

	var filtro bson.M = bson.M{"_id": *medidor}

	var update bson.D = bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "lecturasPendientes", Value: cantidad},
		}},
	}
	_, err := r.collection.UpdateOne(ctx, filtro, update)

	return err

}
