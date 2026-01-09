package service

import (
	"context"
	"errors"
	"prestoBackend/src/core/utils"
	"prestoBackend/src/module/autenticacion/dto"
	"prestoBackend/src/module/usuario/repository"
)

type AutenticacionService struct {
	repository repository.UsuarioRepository
}

func NewAutenticacionService(repository repository.UsuarioRepository) *AutenticacionService {
	return &AutenticacionService{repository: repository}
}

func (controller *AutenticacionService) Autenticacion(dto *dto.AutenticacionDto, ctx context.Context) (string, error) {
	usuario, err := controller.repository.BuscarUsuarioPorUsuario(dto.Usuario, ctx)
	if err != nil {

		return "", errors.New("Credenciales invalidas")
	}
	ok, err := utils.ComparePasswordAndHash(dto.Password, usuario.Password)

	if err != nil || !ok {
		return "", errors.New("Credenciales invalidas")
	}

	token, err := utils.GenraraToken(usuario.ID)
	if err != nil {
		return "", errors.New("Credenciales invalidas")
	}
	return token, nil

}
