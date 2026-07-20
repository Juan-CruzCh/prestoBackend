package repository

import (
	"context"
	"prestoBackend/src/internal/cajaChica/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type CajaChica interface {
	CrearCajaChica(cajaChica *model.CajaChica, ctx context.Context) error
	ListarCajaChica(ctx context.Context) (interface{}, error)
	ActualizarCajaChica(id *bson.ObjectID, cajaChica *model.CajaChica, ctx context.Context) error
	EliminarCajaChica(id *bson.ObjectID, ctx context.Context) error
}

type cajaChica struct {
	collection *mongo.Collection
}

func NewCajaChicaRepository(db *mongo.Database) *cajaChica {
	collection := db.Collection("CajaChica")
	return &cajaChica{collection: collection}
}

func (r *cajaChica) CrearCajaChica(cajaChica *model.CajaChica, ctx context.Context) error {
	return nil
}

func (r *cajaChica) ListarCajaChica(ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (r *cajaChica) ActualizarCajaChica(id *bson.ObjectID, cajaChica *model.CajaChica, ctx context.Context) error {
	return nil
}

func (r *cajaChica) EliminarCajaChica(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
