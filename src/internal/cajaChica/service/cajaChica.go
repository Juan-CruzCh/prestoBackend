package service

import (
	"context"
	"prestoBackend/src/app/common"
	"prestoBackend/src/internal/cajaChica/dto"
	"prestoBackend/src/internal/cajaChica/model"
	"prestoBackend/src/internal/cajaChica/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type CajaChica struct {
	cajaChicaRepository repository.CajaChica
	cliente             *mongo.Client
}

func NewCajaChicaService(cajaChicaRepository repository.CajaChica, cliente *mongo.Client) *CajaChica {
	return &CajaChica{
		cajaChicaRepository: cajaChicaRepository,
		cliente:             cliente,
	}
}

func (s *CajaChica) CrearCajaChica(cajaChica *dto.CajaChicaDto, usuario bson.ObjectID, ctx context.Context) error {
	var data model.CajaChica = model.CajaChica{
		MontoInicial:  cajaChica.MontoInicial,
		MontoActual:   cajaChica.MontoInicial,
		FechaApertura: common.FechaHoraBolivia(),
		Usuario:       usuario,
	}
	err := s.cajaChicaRepository.CrearCajaChica(&data, ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *CajaChica) ListarCajaChica(id *bson.ObjectID, ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (s *CajaChica) ActualizarCajaChica(id *bson.ObjectID, cajaChica *dto.CajaChicaDto, ctx context.Context) error {
	return nil
}

func (s *CajaChica) EliminarCajaChica(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
