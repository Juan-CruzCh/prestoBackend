package controller

import (
	"context"
	"net/http"
	"prestoBackend/src/core/utils"
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
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	validate := validator.New()

	var body dto.BuscadorLecturaDto

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

	resultado, err := controller.service.RepositoryLectura.ListarLectura(&body, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resultado)

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

	if body.LecturaAnterior > body.LecturaActual {
		c.JSON(http.StatusBadRequest, gin.H{"error": "La lectura anterior no debe ser mayor a la lectura actual"})
		return

	}

	resultado, err := controller.service.CrearLectura(&body, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resultado)
}

func (controller *LecturaController) BuscarLecturaPorNumeroMedidor(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	var numeroMedidor string = c.Param("numeroMedidor")
	resultado, err := controller.service.BuscarLecturaPorNumeroMedidor(numeroMedidor, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resultado)
}

func (controller *LecturaController) BuscarLecturasPorClienteMedidor(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	var cliente string = c.Param("cliente")
	IDCliente, err := utils.ValidadIdMongo(cliente)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resultado, err := controller.service.BuscarLecturasPorClienteMedidor(IDCliente, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resultado)
}

func (controller *LecturaController) DetalleLectura(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	var medidor string = c.Param("medidor")
	var lectura string = c.Param("lectura")
	IDmedidor, err := utils.ValidadIdMongo(medidor)
	IDlectura, err := utils.ValidadIdMongo(lectura)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resultado, err := controller.service.DetalleLectura(IDmedidor, IDlectura, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resultado)
}
