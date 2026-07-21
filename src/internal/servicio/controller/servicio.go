package controller

import (
	"net/http"
	"prestoBackend/src/internal/servicio/service"
)

type Servicio struct {
	servicioService *service.Servicio
}

func NewServicioController(servicioService *service.Servicio) *Servicio {
	return &Servicio{
		servicioService: servicioService,
	}
}
func (c *Servicio) CrearServicio(w http.ResponseWriter, r *http.Request) {
}

func (c *Servicio) ListarServicio(w http.ResponseWriter, r *http.Request) {
}

func (c *Servicio) ActualizarServicio(w http.ResponseWriter, r *http.Request) {
}

func (c *Servicio) EliminarServicio(w http.ResponseWriter, r *http.Request) {
}
