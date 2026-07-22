package repository

import (
	"context"
	"fmt"
	"prestoBackend/src/app/common"
	"prestoBackend/src/app/enum"
	"prestoBackend/src/internal/caja/model"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Caja interface {
	VerificarCaja(usuario *bson.ObjectID, ctx context.Context) (bool, error)
	CrearCaja(caja *model.Caja, ctx context.Context) error
	ListarCaja(ctx context.Context) (interface{}, error)
	GurdarPagosEnCaja(caja bson.ObjectID, monto float64, ctx context.Context) error
}

type caja struct {
	collection *mongo.Collection
}

func NewCajaRepository(db *mongo.Database) *caja {
	collection := db.Collection("Caja")
	return &caja{collection: collection}
}

func (r *caja) CrearCaja(caja *model.Caja, ctx context.Context) error {
	caja.Flag = enum.FlagNuevo
	caja.Fecha = common.FechaHoraBolivia()
	cantidad, err := r.contarRegistros(ctx)
	if err != nil {
		return err
	}
	caja.Codigo = "CJ-" + strconv.Itoa(int(*cantidad))
	_, err = r.collection.InsertOne(ctx, caja)
	if err != nil {
		return err
	}
	return nil
}
func (r *caja) VerificarCaja(usuario *bson.ObjectID, ctx context.Context) (*model.Caja, bool, error) {
	hoy := time.Now()
	var filter bson.M = bson.M{
		"usuario": usuario,
		"estado":  enum.Abierto,
	}
	var caja model.Caja
	err := r.collection.FindOne(ctx, filter).Decode(&caja)
	if err != nil {
		if err == mongo.ErrNilDocument {
			return nil, false, fmt.Errorf("Nesesita abrir la caja")
		}
		return nil, false, nil
	}
	if caja.FechaInicio != hoy {
		return nil, false, fmt.Errorf("Existe la caja abierta del dia anterior")
	}
	return &caja, false, nil
}

func (r *caja) GurdarPagosEnCaja(caja bson.ObjectID, monto float64, ctx context.Context) error {
	filter := bson.M{
		"_id":    caja,
		"estado": enum.Abierto,
	}
	update := bson.M{
		"$inc": bson.M{
			"montoPago":  monto,
			"montoTotal": monto,
		},
	}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *caja) contarRegistros(ctx context.Context) (*int64, error) {
	cantidad, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	return &cantidad, nil
}

func (r *caja) ListarCaja(ctx context.Context) (interface{}, error) {
	return nil, nil
}
