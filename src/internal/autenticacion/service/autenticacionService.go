package service

import (
	"context"
	"errors"

	"prestoBackend/src/app/argon"
	"prestoBackend/src/app/jwt"
	"prestoBackend/src/internal/autenticacion/dto"
	"prestoBackend/src/internal/usuario/model"
	"prestoBackend/src/internal/usuario/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type AutenticacionService struct {
	repository repository.UsuarioRepository
}

func NewAutenticacionService(repository repository.UsuarioRepository) *AutenticacionService {
	return &AutenticacionService{repository: repository}
}

func (service *AutenticacionService) Autenticacion(dto *dto.AutenticacionDto, ctx context.Context) (string, error) {
	usuario, err := service.repository.BuscarUsuarioPorUsuario(dto.Usuario, ctx)
	if err != nil {
		return "", errors.New("Credenciales invalidas")
	}
	ok, err := argon.ComparePasswordAndHash(dto.Password, usuario.Password)

	if err != nil || !ok {
		return "", errors.New("Credenciales invalidas")
	}

	token, err := jwt.GenraraToken(usuario.ID)
	if err != nil {
		return "", errors.New("Credenciales invalidas")
	}
	return token, nil

}

func (service *AutenticacionService) BuscarUsuarioPorUsuario(usuarioId *bson.ObjectID, ctx context.Context) (*model.Usuario, error) {
	usuario, err := service.repository.BuscarUsuarioPorUsuarioId(usuarioId, ctx)
	if err != nil {
		return nil, err
	}
	return usuario, nil

}
