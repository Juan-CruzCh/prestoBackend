package router

import (
	"net/http"
	"prestoBackend/src/core/enum"
	"prestoBackend/src/core/middleware"
	"prestoBackend/src/module/lectura/controller"
)

type routerLectura struct {
	controller *controller.LecturaController
	mux        *http.ServeMux
}

func NewLecturaRouter(mux *http.ServeMux, controllerLectura *controller.LecturaController) *routerLectura {
	return &routerLectura{
		controller: controllerLectura,
		mux:        mux,
	}
}

func (r *routerLectura) LecturaRouter() {
	rutaProtegida := middleware.Roles(enum.RolAdministrador, enum.RolLecturador)
	r.mux.Handle("POST /api/lectura/listar", rutaProtegida(http.HandlerFunc(r.controller.ListarLecturas)))
	r.mux.Handle("POST /api/lectura", rutaProtegida(http.HandlerFunc(r.controller.CrearLectura)))
	r.mux.Handle("GET /api/lectura/medidor/{numeroMedidor}", rutaProtegida(http.HandlerFunc(r.controller.BuscarLecturaPorNumeroMedidor)))
	r.mux.Handle("GET /api/lectura/detalle/{medidor}/{lectura}", rutaProtegida(http.HandlerFunc(r.controller.DetalleLectura)))
	r.mux.Handle("GET /api/lectura/medidor/cliente/{cliente}", rutaProtegida(http.HandlerFunc(r.controller.BuscarLecturasPorClienteMedidor)))
	r.mux.Handle("DELETE /api/lectura/{id}", rutaProtegida(http.HandlerFunc(r.controller.EliminarLectura)))
}
