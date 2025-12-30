package dto

type ClienteDto struct {
	Ci              string `json:"ci" validate:"required"`
	Nombre          string `json:"nombre" validate:"required"`
	ApellidoMaterno string `json:"apellidoMaterno"`
	ApellidoPaterno string `json:"apellidoPaterno" validate:"required"`
	Celular         string `json:"celular" validate:"required"`
}

type BucadorClienteDto struct {
	Pagina          int
	Limite          int
	Nombre          string
	Codigo          string
	ApellidoPaterno string
	ApellidoMaterno string
	Ci              string
}
