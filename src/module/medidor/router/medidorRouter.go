package router

import (
	"net/http"
	"prestoBackend/src/module/medidor/controller"
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
	r.mux.HandleFunc("POST /api/medidor", r.controller.CrearMedidor)
	r.mux.HandleFunc("GET /api/medidor", r.controller.ListarMedidorCliente)

	r.mux.HandleFunc("DELETE /api/medidor/{id}", r.controller.EliminarMedidor)
	r.mux.HandleFunc("PATCH /api/medidor/{id}", r.controller.ActualizarMedidor)
	r.mux.HandleFunc("GET /api/medidor/{id}", r.controller.ObtenerMedidorConClientePorId)
}
