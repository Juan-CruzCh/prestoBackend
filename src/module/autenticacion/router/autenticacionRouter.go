package router

import (
	"prestoBackend/src/module/autenticacion/controller"

	"github.com/gin-gonic/gin"
)

func AutenticacionRouter(router *gin.RouterGroup, controller *controller.AutenticacionController) {
	router.POST("autenticacion", controller.Autenticacion)
}
