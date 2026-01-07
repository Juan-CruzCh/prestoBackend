package controller

import (
	"context"
	"net/http"
	"prestoBackend/src/core/utils"
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

func (controller *MedidorController) ListarMedidorCliente(c *gin.Context) {
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
	direccion := c.Query("direccion")
	numeroMedidor := c.Query("numeroMedidor")
	tarifa := c.Query("tarifa")
	estado := c.Query("estado")
	estadoMedidor := c.Query("estadoMedidor")

	var filter dto.BuscadorMedidorClienteDto = dto.BuscadorMedidorClienteDto{
		Pagina:          pagina,
		Limite:          limite,
		Nombre:          nombre,
		Codigo:          codigo,
		ApellidoPaterno: apellidoPaterno,
		ApellidoMaterno: apellidoMaterno,
		Ci:              ci,
		Direccion:       direccion,
		NumeroMedidor:   numeroMedidor,
		Tarifa:          tarifa,
		Estado:          estado,
		EstadoMedidor:   estadoMedidor,
	}

	resultado, err := controller.service.ListarMedidores(&filter, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resultado)
}

func (controller *MedidorController) EliminarMedidor(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	var medidor string = c.Param("id")

	ID, err := utils.ValidadIdMongo(medidor)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultado, err := controller.service.EliminarMedidor(ID, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resultado)
}

func (controller *MedidorController) ActualizarMedidor(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	var id string = c.Param("id")
	ID, err := utils.ValidadIdMongo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()

	var body dto.MedidorDto

	err = c.ShouldBindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = validate.Struct(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resultado, err := controller.service.ActualizarMedidor(ID, &body, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resultado)
}
func (controller *MedidorController) ObtenerMedidorConClientePorId(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	var id string = c.Param("id")
	ID, err := utils.ValidadIdMongo(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resultado, err := controller.service.ObtenerMedidorConClientePorId(ID, ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resultado)

}
