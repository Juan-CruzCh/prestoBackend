package service

import (
	"context"
	"fmt"
	"prestoBackend/src/core/utils"
	"prestoBackend/src/module/lectura/dto"
	lecturaRepository "prestoBackend/src/module/lectura/repository"
	medidorRepository "prestoBackend/src/module/medidor/repository"
	rangoRepository "prestoBackend/src/module/tarifa/repository"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type LecturaService struct {
	RepositoryLectura lecturaRepository.LecturaRepository
	RepositoryRango   rangoRepository.RangoRepository
	RepositoryMedidor medidorRepository.MedidorRepository
}

func NewLecturaService(repositoryLectura lecturaRepository.LecturaRepository, repositoryRango rangoRepository.RangoRepository, RepositoryMedidor medidorRepository.MedidorRepository) *LecturaService {
	return &LecturaService{
		RepositoryLectura: repositoryLectura,
		RepositoryRango:   repositoryRango,
		RepositoryMedidor: RepositoryMedidor,
	}
}
func (s *LecturaService) ListarLectura() {

}

func (s *LecturaService) CrearLectura(lecturaDto *dto.LecturaDto, ctx context.Context) (*mongo.InsertOneResult, error) {
	idMedidor, err := utils.ValidadIdMongo(lecturaDto.Medidor)
	if err != nil {
		return nil, err
	}
	medidor, err := s.RepositoryMedidor.ObtenerMedidor(idMedidor, ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println(medidor)
	return nil, nil
}
