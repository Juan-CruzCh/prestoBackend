package service

import (
	"context"
	"errors"
	"fmt"
	"prestoBackend/src/core/enum"
	"prestoBackend/src/core/utils"
	lecturaModel "prestoBackend/src/module/lectura/model"
	lecturaRepository "prestoBackend/src/module/lectura/repository"
	medidorRepository "prestoBackend/src/module/medidor/repository"
	"prestoBackend/src/module/pago/dto"
	pagoModel "prestoBackend/src/module/pago/model"
	PagoRepository "prestoBackend/src/module/pago/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type PagoService struct {
	PagoRepository        PagoRepository.PagoRepository
	lecturaRepository     lecturaRepository.LecturaRepository
	medidorRepository     medidorRepository.MedidorRepository
	DetallePagoRepository PagoRepository.DetallePagoRepository
	cliente               *mongo.Client
}

func NewPagoService(PagoRepository PagoRepository.PagoRepository,
	lecturaRepository lecturaRepository.LecturaRepository,
	medidorRepository medidorRepository.MedidorRepository,
	DetallePagoRepository PagoRepository.DetallePagoRepository,
	cliente *mongo.Client,
) *PagoService {
	return &PagoService{
		lecturaRepository:     lecturaRepository,
		medidorRepository:     medidorRepository,
		PagoRepository:        PagoRepository,
		DetallePagoRepository: DetallePagoRepository,
		cliente:               cliente,
	}
}

func (service *PagoService) RealizarPago(pagoDto *dto.PagoDto, ctx context.Context) (*bson.ObjectID, error) {
	transaccion, err := service.cliente.StartSession()
	if err != nil {
		return nil, err
	}
	defer transaccion.EndSession(ctx)
	callback := func(ctx context.Context) (any, error) {
		var totalLecturas float64 = 0
		var lecturas []lecturaModel.Lectura = []lecturaModel.Lectura{}
		for _, v := range pagoDto.Lecturas {
			lectura, err := service.lecturaRepository.BuscarLecturaPorId(&v.Lectura, enum.LecturaPendiente, ctx)
			if err != nil {
				return nil, fmt.Errorf("verica tu lectura")
			}

			totalLecturas += lectura.CostoAPagar
			lecturas = append(lecturas, *lectura)
		}
		usuario, err := utils.ValidadIdMongo("67c5f4e9eaa776f45325e80d")
		if err != nil {

			return nil, err
		}

		cantidadPagos, err := service.PagoRepository.CantidadDePagos(ctx)
		if err != nil {

			return nil, err
		}
		var pago pagoModel.Pago = pagoModel.Pago{
			NumeroPago: cantidadPagos,
			Total:      totalLecturas,
			TipoPago:   enum.TipoPagoEfectivo,
			Usuario:    *usuario,
			Flag:       enum.FlagNuevo,
			Fecha:      utils.FechaHoraBolivia(),
		}
		resultado, err := service.PagoRepository.CrearPago(&pago, ctx)
		if err != nil {
			return nil, errors.New("no se pudo registrar el pago")
		}

		var medidor bson.ObjectID

		for _, v := range lecturas {
			medidor = v.Medidor
			var detalle pagoModel.DetallePago = pagoModel.DetallePago{
				Lectura:     v.ID,
				CostoPagado: v.CostoAPagar,
				Flag:        enum.FlagNuevo,
				Fecha:       utils.FechaHoraBolivia(),
				Pago:        resultado.InsertedID.(bson.ObjectID),
			}
			_, err := service.DetallePagoRepository.CrearDetalle(&detalle, ctx)
			if err != nil {

				return nil, err
			}
			_, err = service.lecturaRepository.ActualizarEstadoLectura(&v.ID, enum.LecturaPagado, ctx)
			if err != nil {

				return nil, err
			}
		}
		cantidad, err := service.lecturaRepository.ContarLecturasPorMedidorYEstado(&medidor, enum.LecturaPendiente, ctx)
		if err != nil {

			return nil, err
		}
		err = service.medidorRepository.ActualizaLecturasPendientesMedidor(cantidad, &medidor, ctx)
		if err != nil {

			return nil, err
		}
		return resultado.InsertedID, nil

	}

	resultado, err := transaccion.WithTransaction(ctx, callback)

	if err != nil {
		return nil, err
	}

	ID, ok := resultado.(bson.ObjectID)
	if !ok {
		return nil, err
	}
	return &ID, nil

}
