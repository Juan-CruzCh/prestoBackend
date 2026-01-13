package router

import (
	"net/http"
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

	r.mux.HandleFunc("POST /api/lectura/listar", r.controller.ListarLecturas)
	r.mux.HandleFunc("POST /api/lectura", r.controller.CrearLectura)

	r.mux.HandleFunc("GET /api/lectura/medidor/{numeroMedidor}", r.controller.BuscarLecturaPorNumeroMedidor)

	r.mux.HandleFunc("GET /api/lectura/detalle/{medidor}/{lectura}", r.controller.DetalleLectura)

	r.mux.HandleFunc("GET /api/lectura/medidor/cliente/{cliente}", r.controller.BuscarLecturasPorClienteMedidor)

	r.mux.HandleFunc("DELETE /api/lectura/{id}", r.controller.EliminarLectura)
}
