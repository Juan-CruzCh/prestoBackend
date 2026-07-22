package service

import (
	"context"
	"prestoBackend/src/internal/gasto/dto"
	"prestoBackend/src/internal/gasto/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Gasto struct {
	gastoRepository repository.Gasto
	cliente         *mongo.Client
}

func NewGastoService(gastoRepository repository.Gasto, cliente *mongo.Client) *Gasto {
	return &Gasto{
		gastoRepository: gastoRepository,
		cliente:         cliente,
	}
}
func (s *Gasto) CrearGasto(gasto *dto.GastoDto, ctx context.Context) error {

	return nil
}

func (s *Gasto) ListarGasto(id *bson.ObjectID, ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (s *Gasto) ActualizarGasto(id *bson.ObjectID, gasto *dto.GastoDto, ctx context.Context) error {
	return nil
}

func (s *Gasto) EliminarGasto(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
