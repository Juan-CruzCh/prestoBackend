package router

import (
	"net/http"
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

	r.mux.HandleFunc("GET /api/tarifa/rangos", r.controller.ListarTarifasConRagos)
	r.mux.HandleFunc("GET /api/tarifa", r.controller.ListarTarifas)
	r.mux.HandleFunc("POST /api/tarifa", r.controller.CrearTarifa)

	r.mux.HandleFunc("DELETE /api/tarifa/{id}", r.controller.EliminarTarifa)
	r.mux.HandleFunc("DELETE /api/tarifa/rango/{id}", r.controller.EliminarRango)

}
