package router

import (
	"net/http"
	"prestoBackend/src/internal/caja/controller"
)

func NewCajaRouter(mux *http.ServeMux, controller *controller.Caja) {
	mux.HandleFunc("POST /api/abrir/caja", controller.CrearCaja)
	mux.HandleFunc("GET /api/usuario/caja", controller.ListarCajaPorUsuario)
	mux.HandleFunc("GET /api/caja/listar", controller.ListarCaja)
	mux.HandleFunc("POST /api/caja/cerrar", controller.CerrarCaja)
}
