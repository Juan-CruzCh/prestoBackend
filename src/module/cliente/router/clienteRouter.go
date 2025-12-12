package router

import (
	"prestoBackend/src/module/cliente/controller"

	"github.com/gin-gonic/gin"
)

func ClienteRouter(router *gin.RouterGroup, controller *controller.ClienteController) {
	router.POST("cliente", controller.CrearClienteController)
	router.GET("cliente", controller.ListarClientesController)
	router.PATCH("cliente/:id", controller.ActualizarClienteController)
	router.DELETE("cliente/:id", controller.EliminarClienteController)
}
