package router

import (
	"net/http"
	"prestoBackend/src/module/autenticacion/controller"
)

type routerAutenticacion struct {
	controller *controller.AutenticacionController
	mux        *http.ServeMux
}

func NewAutenticacionRouter(mux *http.ServeMux, controllerAutenticacion *controller.AutenticacionController) *routerAutenticacion {
	return &routerAutenticacion{
		controller: controllerAutenticacion,
		mux:        mux,
	}
}

func (r *routerAutenticacion) AutenticacionRouter() {
	r.mux.HandleFunc("POST /api/autenticacion", r.controller.Autenticacion)
	r.mux.HandleFunc("GET /api/verificar/autenticacion", r.controller.VerificarAutenticacion)

}
