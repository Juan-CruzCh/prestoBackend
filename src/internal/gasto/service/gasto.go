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
	cliente             *mongo.Client
	Validate            *validator.Validate
}

func NewGastoService(gastoRepository repository.Gasto, cajaChicaRepository cajaChicaRepository.CajaChica, cliente *mongo.Client, Validate *validator.Validate) *Gasto {
	return &Gasto{
		gastoRepository:     gastoRepository,
		cliente:             cliente,
		Validate:            Validate,
		cajaChicaRepository: cajaChicaRepository,
	}
}
func (s *Gasto) CrearGasto(gasto *dto.GastoDto, usuario *bson.ObjectID, ctx context.Context) error {

	caja, err := s.cajaChicaRepository.VerificarCajaChica(usuario, ctx)
	if err != nil {
		return err
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
	}
	s.gastoRepository.CrearGasto(&gastoModel, ctx)
	return nil
}

func (s *Gasto) ListarGasto(id *bson.ObjectID, ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (s *Gasto) ActualizarGasto(id *bson.ObjectID, gasto *dto.GastoDto, ctx context.Context) error {
	return nil
}

func (s *Gasto) EliminarGasto(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
