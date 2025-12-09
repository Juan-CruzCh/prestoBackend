package controller

import (
	"context"
	"net/http"
	"prestoBackend/src/module/lectura/dto"
	"prestoBackend/src/module/lectura/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LecturaController struct {
	service *service.LecturaService
}

func NewLecturaController(service *service.LecturaService) *LecturaController {
	return &LecturaController{
		service: service,
	}
}

func (controller *LecturaController) ListarLecturas(c *gin.Context) {
	controller.service.RepositoryLectura.ListarLectura()
}

func (controller *LecturaController) CrearLectura(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	validate := validator.New()

	var body dto.LecturaDto

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

	resultado, err := controller.service.CrearLectura(&body, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resultado)
}
