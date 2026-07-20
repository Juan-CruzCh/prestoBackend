package router

import (
	"net/http"
	"prestoBackend/src/app/enum"
	"prestoBackend/src/app/middleware"
	"prestoBackend/src/internal/autenticacion/controller"
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
	rutaProtegida := middleware.Roles(enum.RolAdministrador, enum.RolLecturador)
	r.mux.Handle("GET /api/verificar/autenticacion", rutaProtegida(http.HandlerFunc(r.controller.VerificarAutenticacion)))
	r.mux.Handle("GET /api/cerrarSession", rutaProtegida(http.HandlerFunc(r.controller.CerrarSession)))

}
