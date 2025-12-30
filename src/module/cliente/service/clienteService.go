package service

import (
	"context"

	"prestoBackend/src/core/coreDto"
	"prestoBackend/src/module/cliente/dto"
	"prestoBackend/src/module/cliente/model"
	"prestoBackend/src/module/cliente/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
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

func (s *ClienteService) CrearCliente(clienteDto *dto.ClienteDto, ctx context.Context) (*model.Cliente, error) {
	var cliente model.Cliente = model.Cliente{
		Ci:              clienteDto.Ci,
		Nombre:          clienteDto.Nombre,
		ApellidoMaterno: clienteDto.ApellidoMaterno,
		ApellidoPaterno: clienteDto.ApellidoPaterno,
		Celular:         clienteDto.Celular,
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

func (s *ClienteService) ActualizarCliente(clienteDto *dto.ClienteDto, ID *bson.ObjectID, ctx context.Context) (*mongo.UpdateResult, error) {
	var cliente model.Cliente = model.Cliente{
		Ci:              clienteDto.Ci,
		Nombre:          clienteDto.Nombre,
		ApellidoMaterno: clienteDto.ApellidoMaterno,
		ApellidoPaterno: clienteDto.ApellidoPaterno,
	}
	resultado, err := s.Repository.ActualizarCliente(&cliente, ID, ctx)
	if err != nil {
		return nil, err
	}
	return resultado, nil
}

func (s *ClienteService) EliminarCliente(ID *bson.ObjectID, ctx context.Context) (*mongo.UpdateResult, error) {
	resultado, err := s.Repository.EliminarCliente(ID, ctx)
	if err != nil {
		return nil, err
	}
	return resultado, nil
}
