package router

import (
	"prestoBackend/src/module/tarifa/controller"

	"github.com/gin-gonic/gin"
)

type routerTarifa struct {
	controller *controller.TarifaController
	router     *gin.RouterGroup
}

func NewTarifaRouter(router *gin.RouterGroup, controllerTarifa *controller.TarifaController) *routerTarifa {
	return &routerTarifa{
		controller: controllerTarifa,
		router:     router,
	}
}

func (r *routerTarifa) TarifaRouter() {
	r.router.GET("tarifa/rangos", r.controller.ListarTarifasConRagos)
	r.router.GET("tarifa", r.controller.ListarTarifas)
	r.router.POST("tarifa", r.controller.CrearTarifa)
	r.router.DELETE("tarifa/:id", r.controller.EliminarTarifa)
	r.router.DELETE("tarifa/rango/:id", r.controller.EliminarRango)
}
