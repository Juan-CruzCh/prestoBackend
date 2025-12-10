package service

import (
	"context"
	"errors"
	"prestoBackend/src/core/enum"
	"prestoBackend/src/core/utils"
	"prestoBackend/src/module/tarifa/dto"
	"prestoBackend/src/module/tarifa/model"
	"prestoBackend/src/module/tarifa/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type TarifaService struct {
	rangoRepository  repository.RangoRepository
	tarifaRepository repository.TarifaRepository
}

func NewTarifaService(rangoRepository repository.RangoRepository, tarifaRepository repository.TarifaRepository) *TarifaService {
	return &TarifaService{
		rangoRepository:  rangoRepository,
		tarifaRepository: tarifaRepository,
	}
}

func (service *TarifaService) ListarTarifasConRagos(ctx context.Context) (*[]bson.M, error) {
	resultado, err := service.tarifaRepository.ListarTarifasConRagos(ctx)
	if err != nil {
		return nil, err
	}
	return resultado, nil
}

func (service *TarifaService) ListarTarifas(ctx context.Context) (*[]bson.M, error) {
	resultado, err := service.tarifaRepository.ListarTarifas(ctx)
	if err != nil {
		return nil, err
	}
	return resultado, nil
}
func (service *TarifaService) CrearTarifa(tarifaDto *dto.TarifaDto, ctx context.Context) (*mongo.InsertOneResult, error) {

	cantidad, err := service.tarifaRepository.VerificarTarifa(tarifaDto.Nombre, ctx)
	if err != nil {
		return nil, err
	}

	if cantidad > 0 {
		return nil, errors.New("la tarifa ya se encuentra registrado")
	}
	var tarifa model.Tarifa = model.Tarifa{
		Nombre: tarifaDto.Nombre,
		Flag:   enum.FlagNuevo,
		Fecha:  utils.FechaHoraBolivia(),
	}
	resultado, err := service.tarifaRepository.CrearTarifa(&tarifa, ctx)

	if err != nil {
		return nil, err
	}
	for _, v := range tarifaDto.Rango {
		var rango model.Rango = model.Rango{
			Rango1: v.Rango1,
			Rango2: v.Rango2,
			Costo:  v.Costo,
			Iva:    v.Iva,
			Tarifa: resultado.InsertedID.(bson.ObjectID),
			Fecha:  utils.FechaHoraBolivia(),
			Flag:   enum.FlagNuevo,
		}
		service.rangoRepository.CrearRango(&rango, ctx)
	}
	return resultado, nil

}
