package router

import (
	"prestoBackend/src/module/medidor/controller"

	"github.com/gin-gonic/gin"
)

func MedidorRouter(router *gin.RouterGroup, controller *controller.MedidorController) {
	router.POST("medidor", controller.CrearMedidor)
	router.GET("medidor", controller.ListarMedidorCliente)
	router.DELETE("medidor/:id", controller.EliminarMedidor)
	router.PATCH("medidor/:id", controller.ActualizarMedidor)
	router.GET("medidor/:id", controller.ObtenerMedidorConClientePorId)

}
