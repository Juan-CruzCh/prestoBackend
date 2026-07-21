package service

import (
	"context"
	"prestoBackend/src/internal/servicio/dto"
	"prestoBackend/src/internal/servicio/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Servicio struct {
	servicioRepository repository.Servicio
	cliente            *mongo.Client
}

func NewServicioService(servicioRepository repository.Servicio, cliente *mongo.Client) *Servicio {
	return &Servicio{
		servicioRepository: servicioRepository,
		cliente:            cliente,
	}
}

func (s *Servicio) CrearServicio(servicio *dto.ServicioDto, ctx context.Context) error {
	return nil
}

func (s *Servicio) ListarServicio(id *bson.ObjectID, ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (s *Servicio) ActualizarServicio(id *bson.ObjectID, servicio *dto.ServicioDto, ctx context.Context) error {
	return nil
}

func (s *Servicio) EliminarServicio(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
