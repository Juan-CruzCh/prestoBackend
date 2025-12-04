package controller

import (
	"context"
	"net/http"
	"prestoBackend/src/module/medidor/dto"
	"prestoBackend/src/module/medidor/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type MedidorController struct {
	service *service.MedidorService
}

func NewMedidorController(service *service.MedidorService) *MedidorController {
	return &MedidorController{
		service: service,
	}

}
func (controller *MedidorController) CrearMedidor(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	validate := validator.New()

	var body dto.MedidorDto

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
	resultado, err := controller.service.CrearMedidor(&body, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resultado)
}
