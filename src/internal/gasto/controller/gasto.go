package controller

import (
	"net/http"
	"prestoBackend/src/internal/gasto/service"
)

type Gasto struct {
	gastoService *service.Gasto
}

func NewGastoController(gastoService *service.Gasto) *Gasto {
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
