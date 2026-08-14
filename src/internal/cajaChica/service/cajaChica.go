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
	f1, f2, err := common.NormalizarRangoDeFechas(cajaChica.FechaInicio, cajaChica.FechaFin)
	if err != nil {
		return err
	}
	var data model.CajaChica = model.CajaChica{
		MontoInicial:  cajaChica.MontoInicial,
		MontoActual:   cajaChica.MontoInicial,
		Usuario:       usuario,
		FechaInicio:   f1,
		FechaFin:      f2,
		MontoRestante: 0,
		CantidadGasto: 0,
		Fecha:         common.FechaHoraBolivia(),
	}
	err = s.cajaChicaRepository.CrearCajaChica(&data, ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *CajaChica) ListarCajaChica(ctx context.Context) (*[]bson.M, error) {
	resultado, err := s.cajaChicaRepository.ListarCajaChica(ctx)
	if err != nil {
		return nil, err
	}
	return resultado, err

}
