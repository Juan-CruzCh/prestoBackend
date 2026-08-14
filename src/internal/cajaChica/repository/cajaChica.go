package repository

import (
	"context"
	"errors"
	"fmt"
	"prestoBackend/src/app/database/aggregation"
	"prestoBackend/src/app/enum"
	"prestoBackend/src/internal/cajaChica/model"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type CajaChica interface {
	CrearCajaChica(cajaChica *model.CajaChica, ctx context.Context) error
	VerificarCajaChica(usuario *bson.ObjectID, ctx context.Context) (*model.CajaChica, error)
	ListarCajaChica(ctx context.Context) (*[]bson.M, error)
	ActulizarMontoCajaChica(caja *bson.ObjectID, monto float64, cantidadGasto int, ctx context.Context) error
}

type cajaChica struct {
	collection *mongo.Collection
}

func NewCajaChicaRepository(db *mongo.Database) *cajaChica {
	collection := db.Collection("CajaChica")
	return &cajaChica{collection: collection}
}

func (r *cajaChica) CrearCajaChica(cajaChica *model.CajaChica, ctx context.Context) error {
	cantidad, err := r.contarRegistros(ctx)
	if err != nil {
		return err
	}
	cajaChica.Estado = enum.Abierto
	cajaChica.Flag = enum.FlagNuevo
	cajaChica.Codigo = "CJCH-" + strconv.Itoa(int(cantidad))
	_, err = r.collection.InsertOne(ctx, cajaChica)
	if err != nil {
		return err
	}
	return nil
}

func (r *cajaChica) VerificarCajaChica(usuario *bson.ObjectID, ctx context.Context) (*model.CajaChica, error) {
	hoy := time.Now()
	var filter bson.M = bson.M{
		"usuario": usuario,
		"estado":  enum.Abierto,
	}
	var caja model.CajaChica
	err := r.collection.FindOne(ctx, filter).Decode(&caja)
	if err != nil {
		if err == mongo.ErrNilDocument {
			return nil, fmt.Errorf("Nesesita abrir la caja")
		}
		return nil, err
	}
	if caja.FechaInicio.Year() != hoy.Year() || caja.FechaInicio.Month() != hoy.Month() || caja.FechaInicio.Day() != hoy.Day() {
		return nil, fmt.Errorf("Existe la caja abierta del mes anterior")
	}
	return &caja, fmt.Errorf("La caja ya se encuetra registrada")
}

func (r *cajaChica) ListarCajaChica(ctx context.Context) (*[]bson.M, error) {
	pipeline := mongo.Pipeline{
		bson.D{
			{
				Key: "$match",
				Value: bson.D{
					{
						Key:   "flag",
						Value: enum.FlagNuevo,
					},
				},
			},
		},
		aggregation.Lookup("Usuario", "usuario", "_id", "usuario"),
		bson.D{
			{
				Key: "$project",
				Value: bson.D{
					{
						Key: "_id", Value: 1,
					},
					{
						Key: "codigo", Value: 1,
					},
					{
						Key: "montoInicial", Value: 1,
					},
					{
						Key: "montoActual", Value: 1,
					},
					{
						Key: "montoRestante", Value: 1,
					},
					{
						Key: "cantidadGasto", Value: 1,
					},
					{
						Key: "fechaInicio", Value: 1,
					},
					{
						Key: "fechaFin", Value: 1,
					},
					{
						Key: "estado", Value: 1,
					},
					{
						Key: "usuario", Value: aggregation.ArrayElemAt("$usuario.usuario", 0),
					},
				},
			},
		},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var cajas []bson.M = []bson.M{}
	for cursor.Next(ctx) {
		var caja bson.M = bson.M{}
		err = cursor.Decode(&caja)
		if err != nil {
			return nil, err
		}
		cajas = append(cajas, caja)
	}
	err = cursor.Err()
	if err != nil {
		return nil, err
	}
	return &cajas, nil
}

func (r *cajaChica) contarRegistros(ctx context.Context) (int64, error) {
	cantidad, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}
	return cantidad + 1, nil
}

func (r *cajaChica) ActulizarMontoCajaChica(caja *bson.ObjectID, monto float64, cantidadGasto int, ctx context.Context) error {
	filter := bson.M{
		"_id":    caja,
		"estado": enum.Abierto,
	}
	update := bson.M{
		"$inc": bson.M{
			"cantidadGasto": cantidadGasto,
			"montoActual":   monto,
		},
	}
	resultado, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if resultado.MatchedCount == 0 {
		return errors.New("la caja no está disponible o ya fue cerrada")
	}
	return nil
}
