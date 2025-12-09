package controller

import (
	"context"
	"net/http"
	"prestoBackend/src/module/pago/dto"
	"prestoBackend/src/module/pago/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type PagoController struct {
	service *service.PagoService
}

func NewPagoController(service *service.PagoService) *PagoController {
	return &PagoController{
		service: service}
}

func (controller *PagoController) RealizarPago(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	validate := validator.New()
	var body dto.PagoDto

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
	resultado, err := controller.service.RealizarPago(&body, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resultado)

}
