package service

import (
	"context"
	"prestoBackend/src/app/argon"
	"prestoBackend/src/app/common"
	"prestoBackend/src/app/enum"
	"prestoBackend/src/internal/usuario/dto"
	"prestoBackend/src/internal/usuario/model"
	"prestoBackend/src/internal/usuario/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
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
	hash, err := argon.EncriptarPassword(*body.Password)
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
		Fecha:           common.FechaHoraBolivia(),
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

func (service *UsuarioService) Eliminar(id *bson.ObjectID, ctx context.Context) (*mongo.UpdateResult, error) {
	resultado, err := service.repository.EliminarUsuario(id, ctx)
	if err != nil {
		return nil, err
	}
	return resultado, nil

}

func (service *UsuarioService) ActualizarUsuario(id *bson.ObjectID, body *dto.UsuarioDto, ctx context.Context) (*mongo.UpdateResult, error) {

	var data model.Usuario = model.Usuario{
		Ci:              body.Ci,
		Nombre:          body.Nombre,
		Celular:         body.Celular,
		ApellidoMaterno: body.ApellidoMaterno,
		ApellidoPaterno: body.ApellidoPaterno,
		Usuario:         body.Usuario,
		Direccion:       body.Direccion,
		Rol:             body.Rol,
	}
	resultado, err := service.repository.ActualizarUsuario(id, &data, ctx)
	if err != nil {
		return nil, err
	}
	return resultado, nil

}
