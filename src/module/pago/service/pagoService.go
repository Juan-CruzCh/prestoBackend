package service

import (
	"context"
	"errors"
	"prestoBackend/src/core/enum"
	"prestoBackend/src/core/utils"
	lecturaRepository "prestoBackend/src/module/lectura/repository"
	medidorRepository "prestoBackend/src/module/medidor/repository"
	"prestoBackend/src/module/pago/dto"
	"prestoBackend/src/module/pago/model"
	PagoRepository "prestoBackend/src/module/pago/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type PagoService struct {
	PagoRepository        PagoRepository.PagoRepository
	lecturaRepository     lecturaRepository.LecturaRepository
	medidorRepository     medidorRepository.MedidorRepository
	DetallePagoRepository PagoRepository.DetallePagoRepository
}

func NewPagoService(PagoRepository PagoRepository.PagoRepository,
	lecturaRepository lecturaRepository.LecturaRepository,
	medidorRepository medidorRepository.MedidorRepository,
	DetallePagoRepository PagoRepository.DetallePagoRepository,
) *PagoService {
	return &PagoService{
		lecturaRepository:     lecturaRepository,
		medidorRepository:     medidorRepository,
		PagoRepository:        PagoRepository,
		DetallePagoRepository: DetallePagoRepository,
	}
}

func (service *PagoService) RealizarPago(pagoDto *dto.PagoDto, ctx context.Context) (*mongo.InsertOneResult, error) {

	var totalLecturas float64 = 0
	for _, v := range pagoDto.Lecturas {
		lectura, err := service.lecturaRepository.BuscarLecturaPorId(&v.Lectura, enum.LecturaPendiente, ctx)
		if err != nil {
			return nil, errors.New("algunas lecturas no existen o ya fueron pagadas ")
		}
		totalLecturas += lectura.CostoAPagar
	}
	usuario, err := utils.ValidadIdMongo("67c5f4e9eaa776f45325e80d")
	if err != nil {
		return nil, err
	}

	cantidadPagos, err := service.PagoRepository.CantidadDePagos(ctx)
	if err != nil {
		return nil, err
	}
	var pagoModel model.Pago = model.Pago{
		NumeroPago: cantidadPagos,
		Total:      totalLecturas,
		TipoPago:   enum.TipoPagoEfectivo,
		Usuario:    *usuario,
		Flag:       enum.FlagNuevo,
		Fecha:      utils.FechaHoraBolivia(),
	}
	resultado, err := service.PagoRepository.CrearPago(&pagoModel, ctx)
	if err != nil {
		return nil, errors.New("no se pudo registrar el pago")
	}

	var medidor bson.ObjectID
	for _, v := range pagoDto.Lecturas {
		lectura, _ := service.lecturaRepository.BuscarLecturaPorId(&v.Lectura, enum.LecturaPendiente, ctx)
		medidor = lectura.Medidor
		var detalle model.DetallePago = model.DetallePago{
			Lectura:     v.Lectura,
			CostoPagado: lectura.CostoAPagar,
			Flag:        enum.FlagNuevo,
			Fecha:       utils.FechaHoraBolivia(),
			Pago:        resultado.InsertedID.(bson.ObjectID),
		}
		service.DetallePagoRepository.CrearDetalle(&detalle, ctx)
		service.lecturaRepository.ActualizarEstadoLectura(&lectura.ID, enum.LecturaPagado, ctx)
	}
	cantidad, err := service.lecturaRepository.ContarLecturasPorMedidorYEstado(&medidor, enum.LecturaPendiente, ctx)
	if err != nil {
		return nil, err
	}
	err = service.medidorRepository.ActualizaLecturasPendientesMedidor(cantidad, &medidor, ctx)
	if err != nil {
		return nil, err
	}
	return resultado, nil

}
