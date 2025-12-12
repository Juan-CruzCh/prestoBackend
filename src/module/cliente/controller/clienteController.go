package controller

import (
	"context"
	"net/http"
	"prestoBackend/src/core/utils"
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

func (controller *ClienteController) ListarClientesController(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	pagina, limite, err := utils.Paginador(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	nombre := c.Query("nombre")
	ci := c.Query("ci")
	codigo := c.Query("codigo")
	apellidoPaterno := c.Query("apellidoPaterno")
	apellidoMaterno := c.Query("apellidoMaterno")
	var filter dto.BucadorClienteDto = dto.BucadorClienteDto{
		Pagina:          pagina,
		Limite:          limite,
		Nombre:          nombre,
		Codigo:          codigo,
		ApellidoPaterno: apellidoPaterno,
		ApellidoMaterno: apellidoMaterno,
		Ci:              ci,
	}

	resultado, err := controller.Service.ListarClientes(filter, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resultado)

}

func (controller *ClienteController) ActualizarClienteController(c *gin.Context) {

}

func (controller *ClienteController) EliminarClienteController(c *gin.Context) {

}
