package router

import (
	"prestoBackend/src/module/usuario/controller"

	"github.com/gin-gonic/gin"
)

type routerUsuario struct {
	controller *controller.UsuarioController
	router     *gin.RouterGroup
}

func NewUsuarioRouter(router *gin.RouterGroup, controllerUsuario *controller.UsuarioController) *routerUsuario {
	return &routerUsuario{
		controller: controllerUsuario,
		router:     router,
	}
}

func (r *routerUsuario) UsuarioRouter() {
	r.router.POST("usuario", r.controller.CrearUsuarios)
	r.router.GET("usuario", r.controller.ListarUsuarios)
	r.router.DELETE("usuario/:id", r.controller.Eliminar)
	r.router.PATCH("usuario/:id", r.controller.ActualizarUsuarios)
}
