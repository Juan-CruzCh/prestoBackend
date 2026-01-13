package router

import (
	"net/http"
	"prestoBackend/src/module/usuario/controller"
)

type routerUsuario struct {
	controller *controller.UsuarioController
	mux        *http.ServeMux
}

func NewUsuarioRouter(mux *http.ServeMux, controllerUsuario *controller.UsuarioController) *routerUsuario {
	return &routerUsuario{
		controller: controllerUsuario,
		mux:        mux,
	}
}

func (r *routerUsuario) UsuarioRouter() {
	r.mux.HandleFunc("POST /api/usuario", r.controller.CrearUsuarios)
	r.mux.HandleFunc("GET /api/usuario", r.controller.ListarUsuarios)
	r.mux.HandleFunc("DELETE /api/usuario/{id}", r.controller.Eliminar)
	r.mux.HandleFunc("PATCH /api/usuario/{id}", r.controller.ActualizarUsuarios)

}
