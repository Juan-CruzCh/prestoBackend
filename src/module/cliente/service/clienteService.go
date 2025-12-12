package service

import (
	"context"
	"errors"

	"prestoBackend/src/core/coreDto"
	"prestoBackend/src/core/enum"
	"prestoBackend/src/core/utils"
	"prestoBackend/src/module/cliente/dto"
	"prestoBackend/src/module/cliente/model"
	"prestoBackend/src/module/cliente/repository"
	"strconv"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ClienteService struct {
	Repository repository.ClienteRepository
}

func NewClienteService(repo repository.ClienteRepository) *ClienteService {
	return &ClienteService{
		Repository: repo,
	}
}

func (s *ClienteService) CrearCliente(clienteDto *dto.ClienteDto, ctx context.Context) (*mongo.InsertOneResult, error) {

	cantidad, err := s.Repository.VerificarClienteCi(clienteDto.Ci, ctx)
	if err != nil {
		return nil, err
	}
	if cantidad > 0 {
		return nil, errors.New("El cliente ya se encuetra registrado")
	}
	cantidad, err = s.Repository.CantidadDocumentosCliente(ctx)
	if err != nil {
		return nil, err
	}
	var cliente model.Cliente = model.Cliente{
		Ci:              clienteDto.Ci,
		Nombre:          clienteDto.Nombre,
		ApellidoMaterno: clienteDto.ApellidoMaterno,
		ApellidoPaterno: clienteDto.ApellidoPaterno,
		Flag:            enum.FlagNuevo,
		Fecha:           utils.FechaHoraBolivia(),
		Codigo:          "C-" + strconv.Itoa(cantidad),
	}

	resultado, err := s.Repository.CrearCliente(&cliente, ctx)

	if err != nil {
		return nil, err
	}
	return resultado, nil
}

func (s *ClienteService) ListarClientes(filter dto.BucadorClienteDto, ctx context.Context) (*coreDto.ResultadoPaginado, error) {
	resultado, err := s.Repository.ListarClientes(filter, ctx)
	if err != nil {
		return nil, err
	}
	return resultado, nil

}

func (s *ClienteService) ActualizarCliente() {

}

func (s *ClienteService) EliminarCliente() {

}
