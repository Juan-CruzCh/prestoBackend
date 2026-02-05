package router

import (
	"net/http"
	"prestoBackend/src/core/enum"
	"prestoBackend/src/core/middleware"
	"prestoBackend/src/module/tarifa/controller"
)

type routerTarifa struct {
	controller *controller.TarifaController
	mux        *http.ServeMux
}

func NewTarifaRouter(mux *http.ServeMux, controllerTarifa *controller.TarifaController) *routerTarifa {
	return &routerTarifa{
		controller: controllerTarifa,
		mux:        mux,
	}
}

func (r *routerTarifa) TarifaRouter() {
	rutaProtegida := middleware.Roles(enum.RolAdministrador)
	r.mux.Handle("GET /api/tarifa/rangos", rutaProtegida(http.HandlerFunc(r.controller.ListarTarifasConRagos)))
	r.mux.Handle("GET /api/tarifa/rangos/{id}", rutaProtegida(http.HandlerFunc(r.controller.ObtenerTarifasRangosId)))
	r.mux.Handle("GET /api/tarifa", rutaProtegida(http.HandlerFunc(r.controller.ListarTarifas)))
	r.mux.Handle("POST /api/tarifa", rutaProtegida(http.HandlerFunc(r.controller.CrearTarifa)))
	r.mux.Handle("PATCH /api/tarifa/{id}", rutaProtegida(http.HandlerFunc(r.controller.EditarTarifa)))
	r.mux.Handle("DELETE /api/tarifa/{id}", rutaProtegida(http.HandlerFunc(r.controller.EliminarTarifa)))

}
