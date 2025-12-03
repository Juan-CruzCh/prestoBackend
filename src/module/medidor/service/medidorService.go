package service

import "prestoBackend/src/module/medidor/repository"

type MedidorService struct {
	repository repository.MedidorRepository
}

func NewMedidoService(repo repository.MedidorRepository) *MedidorService {
	return &MedidorService{
		repository: repo,
	}
}

func (service *MedidorService) ListarMedidores() {
	service.repository.ListarMedidor()
}
