package service

import "prestoBackend/src/module/usuario/repository"

type UsuarioService struct {
	repository repository.UsuarioRepository
}

func NewUsuarioService(repo repository.UsuarioRepository) *UsuarioService {
	return &UsuarioService{
		repository: repo,
	}
}

func (service *UsuarioService) CrearUsuario() {

}
