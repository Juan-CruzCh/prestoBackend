package service

import (
	"context"

	"prestoBackend/src/core/coreDto"
	"prestoBackend/src/module/cliente/dto"
	"prestoBackend/src/module/cliente/model"
	"prestoBackend/src/module/cliente/repository"
	clienteRepository "prestoBackend/src/module/cliente/repository"
	medidorRepository "prestoBackend/src/module/medidor/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ClienteService struct {
	clienteRepository clienteRepository.ClienteRepository
	medidorRepository medidorRepository.MedidorRepository
}

func NewClienteService(clienteRepository repository.ClienteRepository, medidorRepository medidorRepository.MedidorRepository) *ClienteService {
	return &ClienteService{
		clienteRepository: clienteRepository,
		medidorRepository: medidorRepository,
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
	resultado, err := s.clienteRepository.CrearCliente(&cliente, ctx)
	if err != nil {
		return nil, err
	}
	return resultado, nil
}

func (s *ClienteService) ListarClientes(filter dto.BucadorClienteDto, ctx context.Context) (*coreDto.ResultadoPaginado, error) {
	resultado, err := s.clienteRepository.ListarClientes(filter, ctx)
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
		Celular:         clienteDto.Celular,
	}
	resultado, err := s.clienteRepository.ActualizarCliente(&cliente, ID, ctx)
	if err != nil {
		return nil, err
	}
	return resultado, nil
}

func (s *ClienteService) EliminarCliente(ID *bson.ObjectID, ctx context.Context) (*mongo.UpdateResult, error) {
	resultado, err := s.clienteRepository.EliminarCliente(ID, ctx)
	if err != nil {
		return nil, err
	}
	if resultado.ModifiedCount > 0 {
		s.medidorRepository.EliminarMedidoresCliente(ID, ctx)
	}
	return resultado, nil
}
