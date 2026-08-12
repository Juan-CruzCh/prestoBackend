package service

import (
	"context"
	"prestoBackend/src/app/common"
	"prestoBackend/src/app/enum"
	cajaChicaRepository "prestoBackend/src/internal/cajaChica/repository"
	"prestoBackend/src/internal/gasto/dto"
	"prestoBackend/src/internal/gasto/model"
	"prestoBackend/src/internal/gasto/repository"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Gasto struct {
	gastoRepository     repository.Gasto
	cajaChicaRepository cajaChicaRepository.CajaChica
	Cliente             *mongo.Client
	Validate            *validator.Validate
}

func NewGastoService(gastoRepository repository.Gasto, cajaChicaRepository cajaChicaRepository.CajaChica, Cliente *mongo.Client, Validate *validator.Validate) *Gasto {
	return &Gasto{
		gastoRepository:     gastoRepository,
		Cliente:             Cliente,
		Validate:            Validate,
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
		var gastoModel model.Gasto = model.Gasto{
			Descripcion:    gasto.Descripcion,
			Monto:          gasto.Monto,
			CajaChica:      caja.ID,
			Usuario:        *usuario,
			CategoriaGasto: gasto.CategoriaGasto,
			Flag:           enum.FlagNuevo,
			Fecha:          common.FechaHoraBolivia(),
			Comprobante:    "falta",
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

func (s *Gasto) ActualizarGasto(id *bson.ObjectID, gasto *dto.GastoDto, ctx context.Context) error {
	return nil
}

func (s *Gasto) EliminarGasto(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
