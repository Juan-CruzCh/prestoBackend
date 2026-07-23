package controller

import (
	"net/http"
	"prestoBackend/src/internal/servicio/service"

	"github.com/go-playground/validator/v10"
)

type Servicio struct {
	servicioService *service.Servicio
	Validate        *validator.Validate
}

func NewServicioController(servicioService *service.Servicio, Validate *validator.Validate) *Servicio {
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
