package controller

import (
	"context"
	"net/http"
	"prestoBackend/src/module/tarifa/dto"
	"prestoBackend/src/module/tarifa/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TarifaController struct {
	tarifaService *service.TarifaService
}

func NewTarifaController(tarifaService *service.TarifaService) *TarifaController {
	return &TarifaController{
		tarifaService: tarifaService,
	}
}

func (controller *TarifaController) ListarTarifasConRagos(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)

	defer cancel()
	resultado, err := controller.tarifaService.ListarTarifasConRagos(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resultado)
}

func (controller *TarifaController) ListarTarifas(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)

	defer cancel()
	resultado, err := controller.tarifaService.ListarTarifas(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resultado)
}
func (controller *TarifaController) CrearTarifa(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	validate := validator.New()
	var body dto.TarifaDto
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

	resultado, err := controller.tarifaService.CrearTarifa(&body, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resultado)
}

func (controller *TarifaController) EliminarTarifa(c *gin.Context) {
}

func (controller *TarifaController) EliminarRango(c *gin.Context) {
}
