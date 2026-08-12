package router

import (
	"net/http"
	"prestoBackend/src/app/enum"
	"prestoBackend/src/app/middleware"
	"prestoBackend/src/internal/medidor/controller"
)

type routerMedidor struct {
	controller *controller.MedidorController
	mux        *http.ServeMux
}

func NewMedidorRouter(mux *http.ServeMux, controllerMedidor *controller.MedidorController) *routerMedidor {
	return &routerMedidor{
		controller: controllerMedidor,
		mux:        mux,
	}
}

func (r *routerMedidor) MedidorRouter() {
	rutaProtegida := middleware.Roles(enum.RolAdministrador)
	r.mux.Handle("POST /api/medidor", rutaProtegida(http.HandlerFunc(r.controller.CrearMedidor)))
	r.mux.Handle("GET /api/medidor", rutaProtegida(http.HandlerFunc(r.controller.ListarMedidorCliente)))
	r.mux.Handle("DELETE /api/medidor/{id}", rutaProtegida(http.HandlerFunc(r.controller.EliminarMedidor)))
	r.mux.Handle("PATCH /api/medidor/{id}", rutaProtegida(http.HandlerFunc(r.controller.ActualizarMedidor)))
	r.mux.Handle("GET /api/medidor/{id}", rutaProtegida(http.HandlerFunc(r.controller.ObtenerMedidorConClientePorId)))
}
