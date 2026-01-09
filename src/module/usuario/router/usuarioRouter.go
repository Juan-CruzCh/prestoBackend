package router

import (
	"prestoBackend/src/module/usuario/controller"

	"github.com/gin-gonic/gin"
)

func UsuarioRouter(router *gin.RouterGroup, controller *controller.UsuarioController) {
	router.POST("usuario", controller.CrearUsuarios)
	router.GET("usuario", controller.ListarUsuarios)
	router.DELETE("usuario/:id", controller.Eliminar)
	router.PATCH("usuario/:id", controller.ActualizarUsuarios)

}
