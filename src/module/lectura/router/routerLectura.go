package router

import (
	"prestoBackend/src/module/lectura/controller"

	"github.com/gin-gonic/gin"
)

func LecturaRouter(router *gin.RouterGroup, controller *controller.LecturaController) {
	router.POST("lectura/listar", controller.ListarLecturas)
	router.GET("lectura/medidor/:numeroMedidor", controller.BuscarLecturaPorNumeroMedidor)
	router.GET("lectura/detalle/:medidor/:lectura", controller.DetalleLectura)
	router.POST("lectura", controller.CrearLectura)
	router.GET("lectura/medidor/cliente/:cliente", controller.BuscarLecturasPorClienteMedidor)
	router.DELETE("lectura/:id", controller.EliminarLectura)
}
