package repository

import (
	"context"
	"prestoBackend/src/internal/servicio/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Servicio interface {
	CrearServicio(servicio *model.Servicio, ctx context.Context) error
	ListarServicio(ctx context.Context) (interface{}, error)
	ActualizarServicio(id *bson.ObjectID, servicio *model.Servicio, ctx context.Context) error
	EliminarServicio(id *bson.ObjectID, ctx context.Context) error
}

type servicio struct {
	collection *mongo.Collection
}

func NewServicioRepository(db *mongo.Database) *servicio {
	collection := db.Collection("Servicio")
	return &servicio{collection: collection}
}

func (r *servicio) CrearServicio(servicio *model.Servicio, ctx context.Context) error {
	return nil
}

func (r *servicio) ListarServicio(ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (r *servicio) ActualizarServicio(id *bson.ObjectID, servicio *model.Servicio, ctx context.Context) error {
	return nil
}

func (r *servicio) EliminarServicio(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
