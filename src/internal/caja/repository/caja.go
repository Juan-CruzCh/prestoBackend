package repository

import (
	"context"
	"fmt"
	"prestoBackend/src/app/common"
	"prestoBackend/src/app/enum"
	"prestoBackend/src/internal/caja/model"
	"strconv"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Caja interface {
	VerificarCaja(usuario *bson.ObjectID, ctx context.Context) (bool, error)
	CrearCaja(caja *model.Caja, ctx context.Context) error
	ListarCaja(ctx context.Context) (interface{}, error)
	ActualizarCaja(id *bson.ObjectID, caja *model.Caja, ctx context.Context) error
	EliminarCaja(id *bson.ObjectID, ctx context.Context) error
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
func (r *caja) VerificarCaja(usuario *bson.ObjectID, ctx context.Context) (bool, error) {
	var filter bson.M = bson.M{
		"usuario": usuario,
		"estado":  enum.Abierto,
	}
	err := r.collection.FindOne(ctx, filter)
	if err != nil {
		fmt.Println(err.Err())
		return false, nil
	}
	return true, nil
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

func (r *caja) ActualizarCaja(id *bson.ObjectID, caja *model.Caja, ctx context.Context) error {
	return nil
}

func (r *caja) EliminarCaja(id *bson.ObjectID, ctx context.Context) error {
	return nil
}
