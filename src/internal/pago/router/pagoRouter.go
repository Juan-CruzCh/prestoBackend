package router

import (
	"net/http"
	"prestoBackend/src/app/enum"
	"prestoBackend/src/app/middleware"
	"prestoBackend/src/internal/pago/controller"
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
	rutaProtegida := middleware.Roles(enum.RolAdministrador)
	r.mux.Handle("POST /api/pago", rutaProtegida(http.HandlerFunc(r.controller.RealizarPago)))
	r.mux.Handle("GET /api/pago", rutaProtegida(http.HandlerFunc(r.controller.ListarPagos)))
	r.mux.Handle("GET /api/anular/pago/{id}", rutaProtegida(http.HandlerFunc(r.controller.AnularPago)))
	r.mux.Handle("GET /api/pago/detalle/{id}", rutaProtegida(http.HandlerFunc(r.controller.DetallePago)))
}
