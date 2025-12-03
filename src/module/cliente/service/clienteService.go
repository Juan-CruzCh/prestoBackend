package service

import "prestoBackend/src/module/cliente/repository"

type ClienteService struct {
	Repository repository.ClienteRepository
}

func NewClienteService(repo repository.ClienteRepository) ClienteService {
	return ClienteService{
		Repository: repo,
	}
}

func (s *ClienteService) CrearCliente() {
	s.Repository.CrearCliente()

}

func (s *ClienteService) ListarClientes() {

}

func (s *ClienteService) ActualizarCliente() {

}

func (s *ClienteService) EliminarCliente() {

}
