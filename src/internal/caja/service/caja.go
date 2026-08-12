package service

import (
	"context"
	"fmt"
	"prestoBackend/src/app/common"
	"prestoBackend/src/app/enum"
	"prestoBackend/src/internal/caja/dto"
	"prestoBackend/src/internal/caja/model"
	"prestoBackend/src/internal/caja/repository"
	pagoRepository "prestoBackend/src/internal/pago/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Caja struct {
	cajaRepository repository.Caja
	PagoRepository pagoRepository.PagoRepository
	cliente        *mongo.Client
}

func NewCajaService(cajaRepository repository.Caja, PagoRepository pagoRepository.PagoRepository, cliente *mongo.Client) *Caja {
	return &Caja{
		cajaRepository: cajaRepository,
		cliente:        cliente,
		PagoRepository: PagoRepository,
	}
}
func (service *Caja) CrearCaja(caja *dto.CajaDto, usuarioId *bson.ObjectID, ctx context.Context) error {
	_, err := service.cajaRepository.VerificarCaja(usuarioId, ctx)
	if err != nil {
		return err
	}
	cajaMode := model.Caja{
		Usuario:       *usuarioId,
		MontoInicial:  caja.MontoInicial,
		FechaInicio:   common.FechaHoraBolivia(),
		Estado:        enum.Abierto,
		MontoTotal:    caja.MontoInicial,
		MontoPago:     0,
		CantidadPagos: 0,
	}
	err = service.cajaRepository.CrearCaja(&cajaMode, ctx)
	if err != nil {
		return err
	}
	return nil
}
func (service *Caja) ListarCaja(ctx context.Context) (*[]bson.M, error) {
	resultado, err := service.cajaRepository.ListarCaja(ctx)
	if err != nil {
		return nil, err
	}
	return resultado, nil
}

func (service *Caja) ListarCajaPorUsuario(usuario *bson.ObjectID, ctx context.Context) (*model.Caja, error) {
	resultado, err := service.cajaRepository.VerificarCaja(usuario, ctx)
	if err != nil {
		return nil, err
	}
	return resultado, nil
}

func (service *Caja) CerrarCaja(caja *dto.CerrarCajaDto, usuario *bson.ObjectID, ctx context.Context) error {

	cajaUsuario, err := service.cajaRepository.VerificarCaja(usuario, ctx)
	if err != nil {
		return err
	}
	pagos, err := service.PagoRepository.BuscarPagosPorCaja(&cajaUsuario.ID, ctx)
	if err != nil {
		return err
	}
	var totalAcumuladoPagos float64 = 0
	for _, v := range *pagos {
		totalAcumuladoPagos += v.Total
	}

	if caja.MontoTotal > totalAcumuladoPagos {
		return fmt.Errorf("el monto ingresado supera el total acumulado en caja (excedente de %.2f)", caja.MontoTotal-totalAcumuladoPagos)
	}

	if caja.MontoTotal < totalAcumuladoPagos {
		return fmt.Errorf("el monto ingresado es menor al total acumulado en caja (faltante de %.2f)", totalAcumuladoPagos-caja.MontoTotal)
	}
	if caja.MontoTotal != totalAcumuladoPagos {
		return fmt.Errorf("el monto en caja no coincide con el total acumulado (diferencia de %.2f)", caja.MontoTotal-totalAcumuladoPagos)
	}

	err = service.cajaRepository.CerrarCaja(&cajaUsuario.ID, ctx)
	if err != nil {
		return err
	}
	return nil
}
