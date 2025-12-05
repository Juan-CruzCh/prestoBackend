package repository

import (
	"context"
	"errors"
	"prestoBackend/src/core/enum"
	"prestoBackend/src/module/lectura/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type LecturaRepository interface {
	CrearLectura(lectura *model.Lectura, ctx context.Context) (*mongo.InsertOneResult, error)
	ListarLectura()
	ActualizarLectura()
	EliminarLectuta()
}

type lecturaRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewLecturaRepository(db *mongo.Database) LecturaRepository {
	return &lecturaRepository{
		db:         db,
		collection: db.Collection("Lectura"),
	}
}

func (r *lecturaRepository) CrearLectura(lectura *model.Lectura, ctx context.Context) (*mongo.InsertOneResult, error) {

	cantidad, err := r.collection.CountDocuments(ctx, bson.M{"flag": enum.FlagNuevo, "medidor": lectura.Medidor, "mes": lectura.Mes, "gestion": lectura.Gestion})
	if err != nil {
		return nil, err
	}
	if cantidad > 0 {
		return nil, errors.New("la lectura ya se encuetra registrada")
	}
	resultado, err := r.collection.InsertOne(ctx, lectura)
	if err != nil {
		return nil, err
	}
	return resultado, nil

}

func (r *lecturaRepository) ListarLectura() {

}
func (r *lecturaRepository) ActualizarLectura() {

}
func (r *lecturaRepository) EliminarLectuta() {

}
