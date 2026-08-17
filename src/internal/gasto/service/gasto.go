package service

import (
	"context"
	"prestoBackend/src/app/common"
	"prestoBackend/src/app/enum"
	cajaChicaRepository "prestoBackend/src/internal/cajaChica/repository"
	"prestoBackend/src/internal/gasto/dto"
	"prestoBackend/src/internal/gasto/model"
	"prestoBackend/src/internal/gasto/repository"
	"strconv"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Gasto struct {
	gastoRepository     repository.Gasto
	cajaChicaRepository cajaChicaRepository.CajaChica
	Cliente             *mongo.Client
}

func NewGastoService(gastoRepository repository.Gasto, cajaChicaRepository cajaChicaRepository.CajaChica, Cliente *mongo.Client) *Gasto {
	return &Gasto{
		gastoRepository:     gastoRepository,
		Cliente:             Cliente,
		cajaChicaRepository: cajaChicaRepository,
	}
}
func (s *Gasto) CrearGasto(gasto *dto.GastoDto, usuario *bson.ObjectID, ctx context.Context) error {
	session, err := s.Cliente.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)
	_, err = session.WithTransaction(ctx, func(mongoctx context.Context) (any, error) {
		caja, err := s.cajaChicaRepository.VerificarCajaChica(usuario, mongoctx)
		if err != nil {
			return nil, err
		}
		cantidad, err := s.gastoRepository.ContarRegistros(ctx)
		if err != nil {
			return nil, err
		}
		var codigo string = "GST-" + strconv.Itoa(int(cantidad))
		var gastoModel model.Gasto = model.Gasto{
			Codigo:         codigo,
			Descripcion:    gasto.Descripcion,
			Monto:          gasto.Monto,
			CajaChica:      caja.ID,
			Usuario:        *usuario,
			CategoriaGasto: gasto.CategoriaGasto,
			Flag:           enum.FlagNuevo,
			Fecha:          common.FechaHoraBolivia(),
			Comprobante:    gasto.Comprobante,
			Tipo:           enum.EGRESO,
		}
		err = s.gastoRepository.CrearGasto(&gastoModel, mongoctx)
		if err != nil {
			return nil, err
		}
		err = s.cajaChicaRepository.ActulizarMontoCajaChica(&caja.ID, -gasto.Monto, 1, mongoctx)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	if err != nil {
		return err
	}
	return nil
}
func (s *Gasto) ListarGasto(ctx context.Context) (*[]bson.M, error) {
	resultado, err := s.gastoRepository.ListarGasto(ctx)
	if err != nil {
		return nil, err
	}
	return resultado, nil
}

func (s *Gasto) EliminarGasto(id *bson.ObjectID, usuario *bson.ObjectID, ctx context.Context) error {
	session, err := s.Cliente.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)
	_, err = session.WithTransaction(ctx, func(mongoctx context.Context) (any, error) {
		cajaChica, err := s.cajaChicaRepository.VerificarCajaChica(usuario, mongoctx)
		if err != nil {
			return nil, err
		}
		gasto, err := s.gastoRepository.EliminarGasto(id, mongoctx)
		if err != nil {
			return nil, err
		}
		err = s.cajaChicaRepository.ActulizarMontoCajaChica(&cajaChica.ID, gasto.Monto, -1, mongoctx)
		if err != nil {
			return nil, err
		}
		cantidad, err := s.gastoRepository.ContarRegistros(ctx)
		if err != nil {
			return nil, err
		}
		var codigo string = gasto.Codigo + strconv.Itoa(int(cantidad))
		var gastoModel model.Gasto = model.Gasto{
			Codigo:         codigo,
			Descripcion:    gasto.Descripcion,
			Monto:          gasto.Monto,
			CajaChica:      cajaChica.ID,
			Usuario:        *usuario,
			CategoriaGasto: gasto.CategoriaGasto,
			Flag:           enum.FlagNuevo,
			Fecha:          common.FechaHoraBolivia(),
			Tipo:           enum.INGRESO,
		}
		err = s.gastoRepository.CrearGasto(&gastoModel, mongoctx)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	if err != nil {
		return err
	}
	return nil
}
