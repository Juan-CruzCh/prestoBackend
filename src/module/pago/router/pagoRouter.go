package router

import (
	"prestoBackend/src/module/pago/controller"

	"github.com/gin-gonic/gin"
)

type routerPago struct {
	controller *controller.PagoController
	router     *gin.RouterGroup
}

func NewPagoRouter(router *gin.RouterGroup, controllerPago *controller.PagoController) *routerPago {
	return &routerPago{
		controller: controllerPago,
		router:     router,
	}
}

func (r *routerPago) PagoRouter() {
	r.router.POST("pago", r.controller.RealizarPago)
	r.router.GET("pago", r.controller.ListarPagos)
	r.router.GET("pago/detalle/:id", r.controller.DetallePago)
}
