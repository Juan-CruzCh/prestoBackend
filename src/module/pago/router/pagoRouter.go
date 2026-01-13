package router

import (
	"net/http"
	"prestoBackend/src/module/pago/controller"
)

type routerPago struct {
	controller *controller.PagoController
	mux        *http.ServeMux
}

func NewPagoRouter(mux *http.ServeMux, controller *controller.PagoController) *routerPago {
	return &routerPago{
		controller: controller,
		mux:        mux,
	}
}

func (r *routerPago) PagoRouter() {
	r.mux.HandleFunc("POST /api/pago", r.controller.RealizarPago)
	r.mux.HandleFunc("GET /api/pago", r.controller.ListarPagos)
	r.mux.HandleFunc("GET /api/pago/detalle/{id}", r.controller.DetallePago)
}
