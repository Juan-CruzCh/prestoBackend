package router

import (
	"prestoBackend/src/module/lectura/controller"

	"github.com/gin-gonic/gin"
)

type routerLectura struct {
	controller *controller.LecturaController
	router     *gin.RouterGroup
}

func NewLecturaRouter(router *gin.RouterGroup, controllerLectura *controller.LecturaController) *routerLectura {
	return &routerLectura{
		controller: controllerLectura,
		router:     router,
	}
}

func (r *routerLectura) LecturaRouter() {
	r.router.POST("lectura/listar", r.controller.ListarLecturas)
	r.router.GET("lectura/medidor/:numeroMedidor", r.controller.BuscarLecturaPorNumeroMedidor)
	r.router.GET("lectura/detalle/:medidor/:lectura", r.controller.DetalleLectura)
	r.router.POST("lectura", r.controller.CrearLectura)
	r.router.GET("lectura/medidor/cliente/:cliente", r.controller.BuscarLecturasPorClienteMedidor)
	r.router.DELETE("lectura/:id", r.controller.EliminarLectura)
}
