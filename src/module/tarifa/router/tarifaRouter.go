package router

import (
	"prestoBackend/src/module/tarifa/controller"

	"github.com/gin-gonic/gin"
)

func TarifaRouter(router *gin.RouterGroup, controller *controller.TarifaController) {
	router.GET("tarifa/rangos", controller.ListarTarifasConRagos)
	router.GET("tarifa", controller.ListarTarifas)
	router.POST("tarifa", controller.CrearTarifa)
	router.DELETE("tarifa/:id", controller.EliminarTarifa)
	router.DELETE("tarifa/rango/:id", controller.EliminarRango)
}
