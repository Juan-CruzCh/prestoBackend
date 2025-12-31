package router

import (
	"prestoBackend/src/module/medidor/controller"

	"github.com/gin-gonic/gin"
)

func MedidorRouter(router *gin.RouterGroup, controller *controller.MedidorController) {
	router.POST("medidor", controller.CrearMedidor)
	router.GET("medidor", controller.ListarrMedidorCliente)
	router.DELETE("medidor/:id", controller.EliminarMedidor)

}
