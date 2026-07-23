package controller

import (
	"net/http"
	"prestoBackend/src/internal/gasto/service"

	"github.com/go-playground/validator/v10"
)

type Gasto struct {
	gastoService *service.Gasto
	Validate     *validator.Validate
}

func NewGastoController(gastoService *service.Gasto, Validate *validator.Validate) *Gasto {
	return &Gasto{
		gastoService: gastoService,
	}
}
func (c *Gasto) CrearGasto(w http.ResponseWriter, r *http.Request) {
}

func (c *Gasto) ListarGasto(w http.ResponseWriter, r *http.Request) {
}

func (c *Gasto) ActualizarGasto(w http.ResponseWriter, r *http.Request) {
}

func (c *Gasto) EliminarGasto(w http.ResponseWriter, r *http.Request) {
}
