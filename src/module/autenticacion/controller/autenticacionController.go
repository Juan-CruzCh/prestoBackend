package controller

import (
	"prestoBackend/src/module/autenticacion/service"

	"github.com/gin-gonic/gin"
)

type AutenticacionController struct {
	service *service.AutenticacionService
}

func NewAutenticacionController(service *service.AutenticacionService) *AutenticacionController {
	return &AutenticacionController{
		service: service,
	}
}

func (controller *AutenticacionController) Autenticacion(c *gin.Context) {

}
