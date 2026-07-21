package repository

import (
	"context"
	"prestoBackend/src/internal/servicio/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type DetalleServicio interface {
	CrearDetalleServicio(detalleServicio *model.DetalleServicio, ctx context.Context) error
	ListarDetalleServicio(ctx context.Context) (interface{}, error)
	ActualizarDetalleServicio(id *bson.ObjectID, detalleServicio *model.DetalleServicio, ctx context.Context) error
	EliminarDetalleServicio(id *bson.ObjectID, ctx context.Context) error
}

type detalleServicio struct {
	collection *mongo.Collection
}

func NewDetalleServicioRepository(db *mongo.Database) *detalleServicio {
	collection := db.Collection("DetalleServicio")
	return &detalleServicio{collection: collection}
}

func (r *detalleServicio) CrearDetalleServicio(detalleServicio *model.DetalleServicio, ctx context.Context) error {
	return nil
}

func (r *detalleServicio) ListarDetalleServicio(ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (r *detalleServicio) ActualizarDetalleServicio(id *bson.ObjectID, detalleServicio *model.DetalleServicio, ctx context.Context) error {
	return nil
}

func (r *detalleServicio) EliminarDetalleServicio(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
