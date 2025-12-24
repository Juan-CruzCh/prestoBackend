package service

import (
	"prestoBackend/src/module/autenticacion/dto"
	"prestoBackend/src/module/usuario/repository"

	"github.com/gin-gonic/gin"
)

type AutenticacionService struct {
	repository *repository.UsuarioRepository
}

func NewAutenticacionService(repository *repository.UsuarioRepository) *AutenticacionService {
	return &AutenticacionService{repository: repository}
}

func (controller *AutenticacionService) Autenticacion(dto *dto.AutenticacionDto, c *gin.Context) {

}
