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
	NumeroDeLecturaPorMedidor(medidor *bson.ObjectID, ctx context.Context) (int, error)
	CantidadLecturas(ctx context.Context) (int, error)
	ContarLecturasPorMedidorYEstado(medidor *bson.ObjectID, estado enum.EstadoLectura, ctx context.Context) (int, error)
	BuscarLecturaPorId(lectura *bson.ObjectID, estado enum.EstadoLectura, ctx context.Context) (*model.Lectura, error)
	ActualizarEstadoLectura(lectura *bson.ObjectID, estado enum.EstadoLectura, ctx context.Context) (*mongo.UpdateResult, error)
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

func (r *lecturaRepository) NumeroDeLecturaPorMedidor(medidor *bson.ObjectID, ctx context.Context) (int, error) {

	cantidad, err := r.collection.CountDocuments(ctx, bson.M{"flag": enum.FlagNuevo, "medidor": medidor})
	if err != nil {
		return 0, err
	}
	cantidad += 1
	return int(cantidad), nil

}

func (r *lecturaRepository) CantidadLecturas(ctx context.Context) (int, error) {

	cantidad, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}
	cantidad += 1
	return int(cantidad), nil

}

func (r *lecturaRepository) ContarLecturasPorMedidorYEstado(medidor *bson.ObjectID, estado enum.EstadoLectura, ctx context.Context) (int, error) {

	cantidad, err := r.collection.CountDocuments(ctx, bson.M{"flag": enum.FlagNuevo, "estado": estado, "medidor": medidor})
	if err != nil {
		return 0, err
	}
	return int(cantidad), nil

}

func (r *lecturaRepository) BuscarLecturaPorId(lectura *bson.ObjectID, estado enum.EstadoLectura, ctx context.Context) (*model.Lectura, error) {

	var data model.Lectura
	err := r.collection.FindOne(ctx, bson.M{"flag": enum.FlagNuevo, "estado": estado, "_id": lectura}).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil

}

func (r *lecturaRepository) ActualizarEstadoLectura(lectura *bson.ObjectID, estado enum.EstadoLectura, ctx context.Context) (*mongo.UpdateResult, error) {
	var filter bson.M = bson.M{"_id": lectura, "flag": enum.FlagNuevo}
	var update bson.D = bson.D{{Key: "$set", Value: bson.D{{Key: "estado", Value: estado}}}}

	resultado, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return resultado, nil

}
