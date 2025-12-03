package controller

import (
	"prestoBackend/src/module/medidor/service"

	"github.com/gin-gonic/gin"
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
	//ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	//defer cancel()

	controller.service.ListarMedidores()
}
