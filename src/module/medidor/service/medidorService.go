package service

import (
	"context"
	"prestoBackend/src/core/coreDto"
	"prestoBackend/src/core/enum"
	"prestoBackend/src/core/utils"
	"prestoBackend/src/module/medidor/dto"
	"prestoBackend/src/module/medidor/model"
	"prestoBackend/src/module/medidor/repository"
	"strconv"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MedidorService struct {
	repository repository.MedidorRepository
}

func NewMedidoService(repo repository.MedidorRepository) *MedidorService {
	return &MedidorService{
		repository: repo,
	}
}

func (service *MedidorService) ListarMedidores(filter *dto.BuscadorMedidorClienteDto, ctx context.Context) (*coreDto.ResultadoPaginado, error) {
	resultado, err := service.repository.ListarMedidorCliente(filter, ctx)
	if err != nil {

		return nil, err
	}
	return resultado, nil
}

func (service *MedidorService) CrearMedidor(medidorDto *dto.MedidorDto, ctx context.Context) (*mongo.InsertOneResult, error) {
	cantidad, err := service.repository.CantidadMedidor(ctx)
	if err != nil {
		return nil, err
	}
	var medidor model.Medidor = model.Medidor{
		NumeroMedidor:      medidorDto.NumeroMedidor,
		Descripcion:        medidorDto.Descripcion,
		Estado:             enum.MedidorActivo,
		Direccion:          medidorDto.Descripcion,
		FechaInstalacion:   medidorDto.FechaInstalacion,
		Flag:               enum.FlagNuevo,
		Fecha:              utils.FechaHoraBolivia(),
		Codigo:             strconv.Itoa(cantidad),
		Cliente:            medidorDto.Cliente,
		Tarifa:             medidorDto.Tarifa,
		LecturasPendientes: 0,
	}
	resultado, err := service.repository.CrearMedidor(&medidor, ctx)
	if err != nil {
		return nil, err
	}
	return resultado, nil
}
