package dto

type ClienteDto struct {
	Ci              string `json:"ci"`
	Nombre          string `json:"nombre"`
	ApellidoMaterno string `json:"apellidoMaterno"`
	ApellidoPaterno string `json:"apellidoPaterno"`
}
