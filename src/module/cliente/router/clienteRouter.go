package router

import (
	"net/http"
	"prestoBackend/src/core/enum"
	"prestoBackend/src/core/middleware"
	"prestoBackend/src/module/cliente/controller"
)

func NewClienteRouter(mux *http.ServeMux, controller *controller.ClienteController) {
	rutaProtegida := middleware.Roles(enum.RolAdministrador)
	mux.Handle("GET /api/cliente", rutaProtegida(http.HandlerFunc(controller.ListarClientesController)))
	mux.Handle("POST /api/cliente", rutaProtegida(http.HandlerFunc(controller.CrearClienteController)))
	mux.Handle("DELETE /api/cliente/{id}", rutaProtegida(http.HandlerFunc(controller.EliminarClienteController)))
	mux.Handle("PATCH /api/cliente/{id}", rutaProtegida(http.HandlerFunc(controller.ActualizarClienteController)))

}
