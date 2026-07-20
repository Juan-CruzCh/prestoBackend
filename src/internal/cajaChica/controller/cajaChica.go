package controller

import (
	"net/http"
	"prestoBackend/src/internal/cajaChica/service"
)

type CajaChica struct {
	cajaChicaService *service.CajaChica
}

func NewCajaChicaController(cajaChicaService *service.CajaChica) *CajaChica {
	return &CajaChica{
		cajaChicaService: cajaChicaService,
	}
}
func (c *CajaChica) CrearCajaChica(w http.ResponseWriter, r *http.Request) {
}

func (c *CajaChica) ListarCajaChica(w http.ResponseWriter, r *http.Request) {
}

func (c *CajaChica) ActualizarCajaChica(w http.ResponseWriter, r *http.Request) {
}

func (c *CajaChica) EliminarCajaChica(w http.ResponseWriter, r *http.Request) {
}
