package dto

type AutenticacionDto struct {
	Usuario  string `json:"usuario"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}
