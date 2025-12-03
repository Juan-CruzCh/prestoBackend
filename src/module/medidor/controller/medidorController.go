package controller

import "prestoBackend/src/module/medidor/service"

type MedidorController struct {
	service service.MedidorService
}

func NewMedidorController(service service.MedidorService) *MedidorController {
	return &MedidorController{
		service: service,
	}

}
