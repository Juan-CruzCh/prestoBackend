package service

import (
	"context"
	"prestoBackend/src/internal/gasto/dto"
	"prestoBackend/src/internal/gasto/model"
	"prestoBackend/src/internal/gasto/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type CategoriaGasto struct {
	categoriaGastoRepository repository.CategoriaGasto
	cliente                  *mongo.Client
}

func NewCategoriaGastoService(categoriaGastoRepository repository.CategoriaGasto, cliente *mongo.Client) *CategoriaGasto {
	return &CategoriaGasto{
		categoriaGastoRepository: categoriaGastoRepository,
		cliente:                  cliente,
	}
}
func (s *CategoriaGasto) CrearCategoriaGasto(categoriaGasto *dto.CategoriaGastoDto, ctx context.Context) error {
	var categoria model.CategoriaGasto = model.CategoriaGasto{
		Nombre: categoriaGasto.Nombre,
	}
	err := s.categoriaGastoRepository.CrearCategoriaGasto(&categoria, ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *CategoriaGasto) ListarCategoriaGasto(id *bson.ObjectID, ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (s *CategoriaGasto) ActualizarCategoriaGasto(id *bson.ObjectID, categoriaGasto *dto.CategoriaGastoDto, ctx context.Context) error {
	return nil
}

func (s *CategoriaGasto) EliminarCategoriaGasto(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
