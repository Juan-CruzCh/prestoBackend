package router

import (
	"prestoBackend/src/module/tarifa/controller"

	"github.com/gin-gonic/gin"
)

func TarifaRouter(router *gin.RouterGroup, controller *controller.TarifaController) {
	router.GET("tarifa", controller.ListarTarifas)
	router.POST("tarifa", controller.CrearTarifa)
}
