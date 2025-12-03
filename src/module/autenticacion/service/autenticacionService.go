package service

import "prestoBackend/src/module/usuario/repository"

type AutenticacionService struct {
	repository *repository.UsuarioRepository
}

func NewAutenticacionService(repository *repository.UsuarioRepository) *AutenticacionService {
	return &AutenticacionService{repository: repository}
}
