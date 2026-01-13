package router

import (
	"prestoBackend/src/module/cliente/controller"

	"github.com/gin-gonic/gin"
)

type routerCliente struct {
	controller *controller.ClienteController
	router     *gin.RouterGroup
}

func NewClienteRouter(router *gin.RouterGroup, controllerCliente *controller.ClienteController) *routerCliente {
	return &routerCliente{
		controller: controllerCliente,
		router:     router,
	}
}

func (r *routerCliente) ClienteRouter() {
	r.router.POST("cliente", r.controller.CrearClienteController)
	r.router.GET("cliente", r.controller.ListarClientesController)
	r.router.PATCH("cliente/:id", r.controller.ActualizarClienteController)
	r.router.DELETE("cliente/:id", r.controller.EliminarClienteController)
}
