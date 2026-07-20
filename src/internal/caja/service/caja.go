package service

import (
	"context"
	"prestoBackend/src/internal/caja/dto"
	"prestoBackend/src/internal/caja/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Caja struct {
	cajaRepository repository.Caja
	cliente        *mongo.Client
}

func NewCajaService(cajaRepository repository.Caja, cliente *mongo.Client) *Caja {
	return &Caja{
		cajaRepository: cajaRepository,
		cliente:        cliente,
	}
}

func (s *Caja) CrearCaja(caja *dto.CajaDto, ctx context.Context) error {
	return nil
}

func (s *Caja) ListarCaja(id *bson.ObjectID, ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (s *Caja) ActualizarCaja(id *bson.ObjectID, caja *dto.CajaDto, ctx context.Context) error {
	return nil
}

func (s *Caja) EliminarCaja(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
