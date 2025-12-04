package router

import (
	"prestoBackend/src/module/lectura/controller"

	"github.com/gin-gonic/gin"
)

func LecturaRouter(router *gin.RouterGroup, controller *controller.LecturaController) {
	router.GET("lectura", controller.ListarLecturas)
	router.POST("lectura", controller.CrearLectura)
}
