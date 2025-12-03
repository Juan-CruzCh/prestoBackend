package controller

import (
	"prestoBackend/src/module/lectura/service"

	"github.com/gin-gonic/gin"
)

type LecturaController struct {
	service service.LecturaService
}

func NewLecturaController(service *service.LecturaService) *LecturaController {
	return &LecturaController{
		service: *service,
	}
}

func (controller *LecturaController) ListarLecturas(c *gin.Context) {
	controller.service.Repository.ListarLectura()
}
