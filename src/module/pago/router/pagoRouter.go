package router

import (
	"prestoBackend/src/module/pago/controller"

	"github.com/gin-gonic/gin"
)

func PagoRouter(router *gin.RouterGroup, controller *controller.PagoController) {

	router.POST("pago", controller.RealizarPago)
}
