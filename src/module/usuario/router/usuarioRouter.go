package router

import (
	"net/http"
	"prestoBackend/src/core/enum"
	"prestoBackend/src/core/middleware"
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
	rutaProtegida := middleware.Roles(enum.RolAdministrador)
	r.mux.Handle("POST /api/usuario", rutaProtegida(http.HandlerFunc(r.controller.CrearUsuarios)))
	r.mux.Handle("GET /api/usuario", rutaProtegida(http.HandlerFunc(r.controller.ListarUsuarios)))
	r.mux.Handle("DELETE /api/usuario/{id}", rutaProtegida(http.HandlerFunc(r.controller.Eliminar)))
	r.mux.Handle("PATCH /api/usuario/{id}", rutaProtegida(http.HandlerFunc(r.controller.ActualizarUsuarios)))

}
