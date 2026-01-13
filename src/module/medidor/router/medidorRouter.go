package router

import (
	"prestoBackend/src/module/medidor/controller"

	"github.com/gin-gonic/gin"
)

type routerMedidor struct {
	controller *controller.MedidorController
	router     *gin.RouterGroup
}

func NewMedidorRouter(router *gin.RouterGroup, controllerMedidor *controller.MedidorController) *routerMedidor {
	return &routerMedidor{
		controller: controllerMedidor,
		router:     router,
	}
}

func (r *routerMedidor) MedidorRouter() {
	r.router.POST("medidor", r.controller.CrearMedidor)
	r.router.GET("medidor", r.controller.ListarMedidorCliente)
	r.router.DELETE("medidor/:id", r.controller.EliminarMedidor)
	r.router.PATCH("medidor/:id", r.controller.ActualizarMedidor)
	r.router.GET("medidor/:id", r.controller.ObtenerMedidorConClientePorId)
}
