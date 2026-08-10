package service

import (
	"context"
	"errors"
	"fmt"

	"prestoBackend/src/app/common"
	"prestoBackend/src/app/enum"
	CajaRepository "prestoBackend/src/internal/caja/repository"
	lecturaRepository "prestoBackend/src/internal/lectura/repository"
	medidorRepository "prestoBackend/src/internal/medidor/repository"
	"prestoBackend/src/internal/pago/dto"
	pagoModel "prestoBackend/src/internal/pago/model"
	PagoRepository "prestoBackend/src/internal/pago/repository"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type PagoService struct {
	PagoRepository        PagoRepository.PagoRepository
	lecturaRepository     lecturaRepository.LecturaRepository
	medidorRepository     medidorRepository.MedidorRepository
	DetallePagoRepository PagoRepository.DetallePagoRepository
	CajaRepository        CajaRepository.Caja
	Cliente               *mongo.Client
}

func NewPagoService(PagoRepository PagoRepository.PagoRepository,
	lecturaRepository lecturaRepository.LecturaRepository,
	medidorRepository medidorRepository.MedidorRepository,
	DetallePagoRepository PagoRepository.DetallePagoRepository,
	CajaRepository CajaRepository.Caja,
	Cliente *mongo.Client,
) *PagoService {
	return &PagoService{
		lecturaRepository:     lecturaRepository,
		medidorRepository:     medidorRepository,
		PagoRepository:        PagoRepository,
		DetallePagoRepository: DetallePagoRepository,
		Cliente:               Cliente,
		CajaRepository:        CajaRepository,
	}
}

func (service *PagoService) RealizarPago(pagoDto *dto.PagoDto, usuario *bson.ObjectID, ctx context.Context) (*bson.ObjectID, error) {
	caja, err := service.CajaRepository.VerificarCaja(usuario, ctx)
	if err != nil {
		return nil, err
	}
	session, err := service.Cliente.StartSession()

	if err != nil {
		return nil, err
	}

	defer session.EndSession(ctx)

	var pagoId bson.ObjectID
	_, err = session.WithTransaction(ctx, func(mongoCtx context.Context) (any, error) {
		var total float64 = 0
		cantidadPagos, err := service.PagoRepository.CantidadDePagos(mongoCtx)
		if err != nil {

			return nil, err
		}
		var pago pagoModel.Pago = pagoModel.Pago{
			NumeroPago: cantidadPagos,
			Total:      total,
			TipoPago:   enum.TipoPagoEfectivo,
			Usuario:    *usuario,
			Flag:       enum.FlagNuevo,
			Fecha:      common.FechaHoraBolivia(),
			Cliente:    pagoDto.Cliente,
			Medidor:    pagoDto.Medidor,
			Caja:       caja.ID,
		}
		resultado, err := service.PagoRepository.CrearPago(&pago, mongoCtx)
		if err != nil {
			return nil, errors.New("no se pudo registrar el pago")
		}
		idPago, _ := resultado.InsertedID.(bson.ObjectID)
		for _, v := range pagoDto.Lecturas {
			lectura, err := service.lecturaRepository.BuscarLecturaPorId(&v.Lectura, enum.LecturaPendiente, mongoCtx)
			if err != nil {
				return nil, fmt.Errorf("verica tu lectura ", err.Error())
			}

			var detalle pagoModel.DetallePago = pagoModel.DetallePago{
				Lectura:         lectura.ID,
				CostoPagado:     lectura.CostoAPagar,
				Flag:            enum.FlagNuevo,
				Fecha:           common.FechaHoraBolivia(),
				Pago:            idPago,
				Gestion:         lectura.Gestion,
				Mes:             lectura.Mes,
				LecturaActual:   lectura.LecturaActual,
				LecturaAnterior: lectura.LecturaAnterior,
				ConsumoTotal:    lectura.ConsumoTotal,
				CostoAPagar:     lectura.CostoAPagar,
			}
			_, err = service.DetallePagoRepository.CrearDetalle(&detalle, mongoCtx)
			if err != nil {

				return nil, err
			}
			_, err = service.lecturaRepository.ActualizarEstadoLectura(&lectura.ID, enum.LecturaPagado, mongoCtx)
			if err != nil {

				return nil, err
			}
			total += lectura.CostoAPagar

		}
		cantidad, err := service.lecturaRepository.ContarLecturasPorMedidorYEstado(&pagoDto.Medidor, enum.LecturaPendiente, mongoCtx)
		if err != nil {
			return nil, err
		}
		err = service.medidorRepository.ActualizaLecturasPendientesMedidor(cantidad, &pagoDto.Medidor, mongoCtx)
		if err != nil {

			return nil, err
		}
		err = service.PagoRepository.ActualizarMontoPago(&idPago, total, mongoCtx)
		if err != nil {

			return nil, err
		}
		err = service.CajaRepository.GurdarPagosEnCaja(caja.ID, total, 1, mongoCtx)
		if err != nil {

			return nil, err
		}
		pagoId = idPago
		return nil, nil

	})
	if err != nil {
		return nil, err
	}

	return &pagoId, nil

}

func (service *PagoService) DetallePago(idPago *bson.ObjectID, ctx context.Context) (*map[string]interface{}, error) {

	pago, err := service.PagoRepository.BuscarPagoId(idPago, ctx)
	if err != nil {
		return nil, err
	}

	detallePago, err := service.PagoRepository.DetallePago(&pago.ID, ctx)
	if err != nil {
		return nil, err
	}

	historial, err := service.lecturaRepository.HistorialLecturaMedidor(&pago.Medidor, ctx)
	if err != nil {
		return nil, err
	}

	data := map[string]interface{}{
		"detallePago": detallePago,
		"historial":   historial,
		"gestion":     time.Now().Year(),
	}

	return &data, nil
}

func (service *PagoService) ListarPagos(filter *dto.BuscardorPagoDto, ctx context.Context) (*map[string]interface{}, error) {
	resultado, err := service.PagoRepository.ListarPagos(filter, ctx)
	if err != nil {
		return nil, err
	}

	return resultado, nil
}
