package controller

import (
	"context"
	"net/http"
	"prestoBackend/src/core/utils"
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

func (controller *PagoController) DetallePago(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	var idPago string = c.Param("id")
	ID, err := utils.ValidadIdMongo(idPago)
	resultado, err := controller.service.DetallePago(ID, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resultado)

}

func (controller *PagoController) ListarPagos(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()
	pagina, limite, err := utils.Paginador(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var filter dto.BuscardorPagoDto = dto.BuscardorPagoDto{
		CodigoCliente:   c.Param("CodigoCliente"),
		Ci:              c.Param("ci"),
		Nombre:          c.Param("nombre"),
		ApellidoMaterno: c.Param("apellidoMaterno"),
		ApellidoPaterno: c.Param("apellidoPaterno"),
		NumeroMedidor:   c.Param("numeroMedidor"),
		FechaInicio:     c.Param("fechaInicio"),
		FechaFin:        c.Param("fechaFin"),
		Pagina:          pagina,
		Limite:          limite,
	}

	resultado, err := controller.service.ListarPagos(&filter, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resultado)

}
