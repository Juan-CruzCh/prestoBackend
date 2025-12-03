package controller

import (
	"prestoBackend/src/module/cliente/service"

	"github.com/gin-gonic/gin"
)

type ClienteController struct {
	Service service.ClienteService
}

func NewClienteController(s service.ClienteService) *ClienteController {
	return &ClienteController{
		Service: s,
	}
}

func (ctl *ClienteController) CrearClienteController(c *gin.Context) {
	ctl.Service.CrearCliente()
}

func ListarClientesController(c *gin.Context) {

}

func ActualizarClienteController(c *gin.Context) {

}

func EliminarClienteController(c *gin.Context) {

}
