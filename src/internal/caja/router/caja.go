package router

import (
	"net/http"
	"prestoBackend/src/internal/caja/controller"
)

func NewCajaRouter(mux *http.ServeMux, controller *controller.Caja) {
	mux.HandleFunc("POST /api/abrir/caja", controller.CrearCaja)
	mux.HandleFunc("GET /api/usuario/caja", controller.ListarCajaPorUsuario)

}
