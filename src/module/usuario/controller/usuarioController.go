package controller

import (
	"prestoBackend/src/module/usuario/service"

	"github.com/gin-gonic/gin"
)

type UsuarioController struct {
	service *service.UsuarioService
}

func NewUsuarioController(service *service.UsuarioService) *UsuarioController {
	return &UsuarioController{
		service: service,
	}
}

func (controller *UsuarioController) ListarUsuarios(c *gin.Context) {

}

func (controller *UsuarioController) CrearUsuarios(c *gin.Context) {

}
