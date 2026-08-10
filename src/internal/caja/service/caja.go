package service

import (
	"context"
	"prestoBackend/src/app/common"
	"prestoBackend/src/app/enum"
	"prestoBackend/src/internal/caja/dto"
	"prestoBackend/src/internal/caja/model"
	"prestoBackend/src/internal/caja/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Caja struct {
	cajaRepository repository.Caja
	cliente        *mongo.Client
}

func NewCajaService(cajaRepository repository.Caja, cliente *mongo.Client) *Caja {
	return &Caja{
		cajaRepository: cajaRepository,
		cliente:        cliente,
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
func (service *Caja) ListarCaja(id *bson.ObjectID, ctx context.Context) (interface{}, error) {
	return nil, nil
}

func (service *Caja) ListarCajaPorUsuario(usuario *bson.ObjectID, ctx context.Context) (*model.Caja, error) {
	resultado, err := service.cajaRepository.ListarCajaPorUsuario(usuario, ctx)
	if err != nil {
		return nil, err
	}
	return resultado, nil
}
