package controller

import (
	"context"
	"net/http"
	"prestoBackend/src/module/cliente/dto"
	"prestoBackend/src/module/cliente/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ClienteController struct {
	Service *service.ClienteService
}

func NewClienteController(s *service.ClienteService) *ClienteController {
	return &ClienteController{
		Service: s,
	}
}

func (ctl *ClienteController) CrearClienteController(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	validate := validator.New()
	var body dto.ClienteDto
	err := c.ShouldBindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = validate.Struct(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resultado, err := ctl.Service.CrearCliente(&body, ctx)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resultado)

}

func ListarClientesController(c *gin.Context) {

}

func ActualizarClienteController(c *gin.Context) {

}

func EliminarClienteController(c *gin.Context) {

}
