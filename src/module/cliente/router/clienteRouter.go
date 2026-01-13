package router

import (
	"net/http"
	"prestoBackend/src/module/cliente/controller"
)

func NewClienteRouter(mux *http.ServeMux, controller *controller.ClienteController) {
	mux.HandleFunc("GET /api/cliente", controller.ListarClientesController)
	mux.HandleFunc("POST /api/cliente", controller.CrearClienteController)
	mux.HandleFunc("DELETE /api/cliente/{id}", controller.EliminarClienteController)
	mux.HandleFunc("PATCH /api/cliente/{id}", controller.EliminarClienteController)

}
