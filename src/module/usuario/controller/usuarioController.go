package controller

import (
	"context"
	"net/http"
	"prestoBackend/src/module/usuario/dto"
	"prestoBackend/src/module/usuario/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UsuarioController struct {
	service *service.UsuarioService
}

func NewUsuarioController(service *service.UsuarioService) *UsuarioController {
	return &UsuarioController{
		service: service,
	}
}

func (controller *UsuarioController) CrearUsuarios(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	validate := validator.New()
	var body dto.UsuarioDto

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = validate.Struct(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resultado, err := controller.service.CrearUsuario(&body, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resultado)

}

func (controller *UsuarioController) ListarUsuarios(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	resultado, err := controller.service.ListarUsuarios(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resultado)

}
