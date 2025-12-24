package service

import (
	"context"
	"prestoBackend/src/core/enum"
	"prestoBackend/src/core/utils"
	"prestoBackend/src/module/usuario/dto"
	"prestoBackend/src/module/usuario/model"
	"prestoBackend/src/module/usuario/repository"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UsuarioService struct {
	repository repository.UsuarioRepository
}

func NewUsuarioService(repo repository.UsuarioRepository) *UsuarioService {
	return &UsuarioService{
		repository: repo,
	}
}

func (service *UsuarioService) CrearUsuario(body *dto.UsuarioDto, ctx context.Context) (*mongo.InsertOneResult, error) {
	hash, err := utils.EncriptarPassword(body.Password)
	var data model.Usuario = model.Usuario{
		Ci:              body.Ci,
		Nombre:          body.Nombre,
		Celular:         body.Celular,
		ApellidoMaterno: body.ApellidoMaterno,
		ApellidoPaterno: body.ApellidoPaterno,
		Usuario:         body.Usuario,
		Password:        hash,
		Direccion:       body.Direccion,
		Flag:            enum.FlagNuevo,
		Rol:             body.Rol,
		Fecha:           utils.FechaHoraBolivia(),
	}
	resultado, err := service.repository.CrearUsuario(&data, ctx)
	if err != nil {
		return nil, err
	}
	return resultado, nil

}

func (service *UsuarioService) ListarUsuarios(ctx context.Context) (*[]model.Usuario, error) {

	resultado, err := service.repository.ListarUsuario(ctx)
	if err != nil {
		return nil, err
	}
	return resultado, nil

}
