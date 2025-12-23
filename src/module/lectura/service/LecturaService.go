package service

import (
	"context"
	"errors"
	"fmt"
	"prestoBackend/src/core/enum"
	"prestoBackend/src/core/utils"
	"prestoBackend/src/module/lectura/dto"
	"prestoBackend/src/module/lectura/model"
	lecturaRepository "prestoBackend/src/module/lectura/repository"
	medidorRepository "prestoBackend/src/module/medidor/repository"
	rangoRepository "prestoBackend/src/module/tarifa/repository"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
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
func (s *LecturaService) ListarLectura(filter *dto.BuscadorLecturaDto, ctx context.Context) (*[]bson.M, error) {
	resultado, err := s.RepositoryLectura.ListarLectura(filter, ctx)
	if err != nil {
		return nil, err

	}
	return resultado, nil
}

func (s *LecturaService) CrearLectura(lecturaDto *dto.LecturaDto, ctx context.Context) (*mongo.InsertOneResult, error) {
	fechaActual := time.Now()
	fechaVencimiento := fechaActual.AddDate(0, 3, 0)
	fmt.Println(lecturaDto)
	if lecturaDto.LecturaActual < lecturaDto.LecturaAnterior {
		return nil, fmt.Errorf("Verifica tu lectura ingresada")
	}
	var consumoAgua int = lecturaDto.LecturaActual - lecturaDto.LecturaAnterior
	medidor, err := s.RepositoryMedidor.ObtenerMedidor(&lecturaDto.Medidor, ctx)
	if err != nil {
		return nil, err
	}
	total, err := s.calcularTarifa(medidor.Tarifa, consumoAgua, ctx)
	if err != nil {
		return nil, err
	}
	numeroLectura, err := s.RepositoryLectura.NumeroDeLecturaPorMedidor(&medidor.ID, ctx)
	if err != nil {
		return nil, err
	}

	cantidadLecturas, err := s.RepositoryLectura.CantidadLecturas(ctx)
	if err != nil {
		return nil, err
	}

	usuario, err := utils.ValidadIdMongo("67c5f4e9eaa776f45325e80d")
	if err != nil {
		return nil, err
	}
	var lectura model.Lectura = model.Lectura{
		Codigo:           "LCT-" + strconv.Itoa(cantidadLecturas),
		NumeroLectura:    numeroLectura,
		Mes:              lecturaDto.Mes,
		LecturaActual:    lecturaDto.LecturaActual,
		LecturaAnterior:  lecturaDto.LecturaAnterior,
		ConsumoTotal:     consumoAgua,
		CostoAPagar:      total,
		Gestion:          lecturaDto.Gestion,
		Estado:           enum.LecturaPendiente,
		Medidor:          medidor.ID,
		Usuario:          *usuario,
		Flag:             enum.FlagNuevo,
		Fecha:            utils.FechaHoraBolivia(),
		FechaVencimiento: fechaVencimiento,
	}
	resultado, err := s.RepositoryLectura.CrearLectura(&lectura, ctx)

	if err != nil {
		return nil, err

	}
	cantidad, _ := s.RepositoryLectura.ContarLecturasPorMedidorYEstado(&medidor.ID, enum.LecturaPendiente, ctx)
	s.RepositoryMedidor.ActualizaLecturasPendientesMedidor(cantidad, &medidor.ID, ctx)
	return resultado, nil
}

func (s *LecturaService) calcularTarifa(tarifa bson.ObjectID, consumoAgua int, ctx context.Context) (float64, error) {

	rangos, err := s.RepositoryRango.ListarRangoPorTarifa(&tarifa, ctx)
	if err != nil {
		return 0, err
	}
	var total float64 = 0

	if consumoAgua <= 0 {
		consumoAgua = 1
	}
	for _, v := range *rangos {
		var iva float64 = v.Iva / 100
		if consumoAgua >= v.Rango1 && consumoAgua <= v.Rango2 {
			var costo float64 = float64(consumoAgua) * v.Costo
			var constoIva float64 = costo * iva
			total = utils.RoundFloat(costo+constoIva, 2)
			break
		}

	}
	return total, nil
}

func (service *LecturaService) BuscarLecturaPorNumeroMedidor(numeroMedidor string, ctx context.Context) (any, error) {
	medidor, err := service.RepositoryMedidor.BuscarMedidorPorNumeroMedidor(numeroMedidor, ctx)

	if err != nil {
		return nil, err
	}

	if len(medidor) <= 0 {
		return nil, fmt.Errorf("El medidor con numero " + numeroMedidor + " no existe")
	}
	data := map[string]interface{}{
		"medidor":       medidor[0].ID,
		"numeroMedidor": medidor[0].NumeroMedidor,
		"estado":        medidor[0].Estado,
		"nombre":        medidor[0].Nombre,
		"apellidos":     medidor[0].ApellidoPaterno + " " + medidor[0].ApellidoMaterno,
	}
	lectura, err := service.RepositoryLectura.UltimaLecturaMedidor(&medidor[0].ID, ctx)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			data["lecturaActual"] = 0
			data["lecturaAnterior"] = 0
			return data, nil
		}
		return nil, err
	}

	data["lecturaActual"] = lectura.LecturaActual
	data["lecturaAnterior"] = lectura.LecturaAnterior

	return data, nil
}

func (s *LecturaService) BuscarLecturasPorClienteMedidor(cliente *bson.ObjectID, ctx context.Context) (*[]interface{}, error) {
	medidor, err := s.RepositoryMedidor.BuscarMedidorCliente(cliente, ctx)
	if err != nil {
		return nil, err
	}
	var resultado []interface{} = []interface{}{}
	for _, v := range medidor {
		lecturas, err := s.RepositoryLectura.LecturasPorMedidor(&v.ID, ctx)
		if err != nil {
			return nil, err
		}
		medidores := map[string]interface{}{
			"_id":           v.ID,
			"numeroMedidor": v.NumeroMedidor,
			"direccion":     v.Direccion,
			"estado":        v.Estado,
			"lecturas":      lecturas,
		}

		resultado = append(resultado, medidores)

	}

	return &resultado, nil

}
