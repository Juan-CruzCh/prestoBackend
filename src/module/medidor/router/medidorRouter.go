package router

import (
	"prestoBackend/src/module/medidor/controller"

	"github.com/gin-gonic/gin"
)

func MedidorRouter(router *gin.RouterGroup, controller *controller.MedidorController) {
	router.GET("medidor", controller.CrearMedidor)
}
