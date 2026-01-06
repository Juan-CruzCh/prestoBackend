package dto

import "prestoBackend/src/core/enum"

type UsuarioDto struct {
	Ci              string    `json:"ci"  validate:"required"`
	Nombre          string    `json:"nombre"  validate:"required"`
	Celular         string    `json:"celular"  validate:"required"`
	ApellidoMaterno string    `json:"apellidoMaterno"  `
	ApellidoPaterno string    `json:"apellidoPaterno" validate:"required"`
	Usuario         string    `json:"usuario"  validate:"required"`
	Password        *string   `json:"password,omitempty" validate:"omitempty,min=6"`
	Direccion       string    `json:"direccion"  validate:"required"`
	Rol             enum.RolE `json:"rol"  validate:"required"`
}
