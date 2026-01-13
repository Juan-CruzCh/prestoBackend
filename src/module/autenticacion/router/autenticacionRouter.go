package router

import (
	"prestoBackend/src/module/autenticacion/controller"

	"github.com/gin-gonic/gin"
)

type routerAutenticacion struct {
	controller *controller.AutenticacionController
	router     *gin.RouterGroup
}

func NewAutenticacionRouter(router *gin.RouterGroup, controllerAutenticacion *controller.AutenticacionController) *routerAutenticacion {
	return &routerAutenticacion{
		controller: controllerAutenticacion,
		router:     router,
	}
}

func (r *routerAutenticacion) AutenticacionRouter() {
	r.router.POST("autenticacion", r.controller.Autenticacion)
}
