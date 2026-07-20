package controller

import (
	"net/http"
	"prestoBackend/src/internal/caja/service"
)

type Caja struct {
	cajaService *service.Caja
}

func NewCajaController(cajaService *service.Caja) *Caja {
	return &Caja{
		cajaService: cajaService,
	}
}
func (c *Caja) CrearCaja(w http.ResponseWriter, r *http.Request) {
}

func (c *Caja) ListarCaja(w http.ResponseWriter, r *http.Request) {
}

func (c *Caja) ActualizarCaja(w http.ResponseWriter, r *http.Request) {
}

func (c *Caja) EliminarCaja(w http.ResponseWriter, r *http.Request) {
}
