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
	if err.Error() != "Nesesita abrir la caja" {
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
func (service *Caja) VerCajaPorUsuarioConSusPagos(usuario *bson.ObjectID, ctx context.Context) (*map[string]interface{}, error) {
	cajaUsuario, err := service.cajaRepository.BuscarCajaPorUsuario(usuario, ctx)
	if err != nil {
		return nil, err
	}
	var resultado map[string]interface{} = map[string]interface{}{}
	if len(cajaUsuario) > 0 {
		cajaId := cajaUsuario[0]["_id"]
		id, ok := cajaId.(bson.ObjectID)
		if !ok {
			return nil, fmt.Errorf("el _id no es un ObjectID válido")
		}
		pagos, err := service.PagoRepository.BuscarPagosPorCaja(&id, ctx)
		if err != nil {
			return nil, err
		}
		resultado = map[string]interface{}{
			"cajaId":        cajaUsuario[0]["_id"],
			"codigo":        cajaUsuario[0]["codigo"],
			"montoInicial":  cajaUsuario[0]["montoInicial"],
			"montoPago":     cajaUsuario[0]["montoPago"],
			"montoTotal":    cajaUsuario[0]["montoTotal"],
			"cantidadPagos": cajaUsuario[0]["cantidadPagos"],
			"fechaInicio":   cajaUsuario[0]["fechaInicio"],
			"fechaFin":      cajaUsuario[0]["fechaFin"],
			"estado":        cajaUsuario[0]["estado"],
			"usuario":       cajaUsuario[0]["usuario"],
			"pagos":         pagos,
		}
	}
	return &resultado, nil
}

func (service *Caja) CerrarCaja(caja *dto.CerrarCajaDto, usuario *bson.ObjectID, ctx context.Context) error {
	cajaUsuario, err := service.cajaRepository.BuscarCajaPorUsuario(usuario, ctx)
	if err != nil {
		return err
	}
	cajaId := cajaUsuario[0]["_id"]
	id, ok := cajaId.(bson.ObjectID)
	if !ok {
		return fmt.Errorf("el _id no es un ObjectID válido")
	}
	montoInicialCaja := cajaUsuario[0]["montoInicial"]
	montoInicial, ok := montoInicialCaja.(float64)
	if !ok {
		return fmt.Errorf("Error en el parseo de monto inicial")
	}
	pagos, err := service.PagoRepository.BuscarPagosPorCaja(&id, ctx)
	if err != nil {
		return err
	}
	var totalAcumuladoPagos float64 = montoInicial
	for _, v := range pagos {
		total, ok := v["total"].(float64)
		if !ok {
			return fmt.Errorf("el campo total no es float64")
		}
		totalAcumuladoPagos += total
	}
	if caja.MontoTotal > totalAcumuladoPagos {
		return fmt.Errorf("el monto ingresado supera el total acumulado en caja (excedente de %.2f)", caja.MontoTotal-totalAcumuladoPagos)
	}

	if caja.MontoTotal < totalAcumuladoPagos {
		return fmt.Errorf("el monto ingresado es menor al total acumulado en caja (faltante de %.2f)", totalAcumuladoPagos-caja.MontoTotal)
	}
	err = service.cajaRepository.CerrarCaja(&id, ctx)
	if err != nil {
		return err
	}
	return nil
}
