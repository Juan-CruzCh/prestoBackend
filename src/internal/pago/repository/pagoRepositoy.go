package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"prestoBackend/src/app/common"
	"prestoBackend/src/app/database/aggregation"
	"prestoBackend/src/app/enum"
	"prestoBackend/src/internal/pago/dto"
	"prestoBackend/src/internal/pago/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type PagoRepository interface {
	CrearPago(pago *model.Pago, cxt context.Context) (*mongo.InsertOneResult, error)
	CantidadDePagos(cxt context.Context) (int, error)
	DetallePago(idPago *bson.ObjectID, ctx context.Context) (*bson.M, error)
	BuscarPagoId(idPago *bson.ObjectID, cxt context.Context) (model.Pago, error)
	ListarPagos(filter *dto.BuscardorPagoDto, ctx context.Context) (*map[string]interface{}, error)
	ActualizarMontoPago(pago *bson.ObjectID, total float64, cxt context.Context) error
	AnularPago(idPago *bson.ObjectID, cxt context.Context) (*model.Pago, error)
	BuscarPagosPorCaja(caja *bson.ObjectID, cxt context.Context) (*[]model.Pago, error)
}

type pagoRepository struct {
	bd         *mongo.Database
	collection *mongo.Collection
}

func NewPagoRepository(bd *mongo.Database) PagoRepository {
	return &pagoRepository{
		bd:         bd,
		collection: bd.Collection("Pago"),
	}

}

func (repo *pagoRepository) CrearPago(pago *model.Pago, cxt context.Context) (*mongo.InsertOneResult, error) {
	resultado, err := repo.collection.InsertOne(cxt, pago)
	if err != nil {
		return nil, errors.New("ocurrio un error al realizar el pago")
	}
	return resultado, nil

}
func (repo *pagoRepository) ActualizarMontoPago(pago *bson.ObjectID, total float64, cxt context.Context) error {
	_, err := repo.collection.UpdateOne(cxt, bson.M{"_id": pago}, bson.M{"$set": bson.M{"total": total}})
	if err != nil {
		return errors.New("ocurrio un error al actulizar el monto")
	}
	return nil

}

func (repo *pagoRepository) CantidadDePagos(cxt context.Context) (int, error) {
	cantidad, err := repo.collection.CountDocuments(cxt, bson.M{})
	if err != nil {
		return 0, errors.New("ocurrio un error al realizar el pag")
	}
	cantidad += 1
	return int(cantidad), nil

}

func (repo *pagoRepository) BuscarPagoId(idPago *bson.ObjectID, cxt context.Context) (model.Pago, error) {
	var data model.Pago
	err := repo.collection.FindOne(cxt, bson.M{"_id": idPago, "flag": enum.FlagNuevo}).Decode(&data)
	if err != nil {
		return model.Pago{}, err
	}
	return data, nil

}

func (repo *pagoRepository) DetallePago(idPago *bson.ObjectID, ctx context.Context) (*bson.M, error) {
	var pipepine mongo.Pipeline = mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{
				{
					Key: "_id", Value: idPago,
				},
				{
					Key: "flag", Value: enum.FlagNuevo,
				},
			}},
		},

		aggregation.Lookup("Cliente", "cliente", "_id", "cliente"),
		aggregation.Lookup("Medidor", "medidor", "_id", "medidor"),
		aggregation.Lookup("DetallePago", "_id", "pago", "detallePago"),

		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "numeroPago", Value: 1},
				{Key: "total", Value: 1},
				{Key: "fecha", Value: 1},
				{Key: "numeroMedidor", Value: aggregation.ArrayElemAt("$medidor.numeroMedidor", 0)},
				{Key: "nombre", Value: aggregation.ArrayElemAt("$cliente.nombre", 0)},
				{Key: "apellidoPaterno", Value: aggregation.ArrayElemAt("$cliente.apellidoPaterno", 0)},
				{Key: "apellidoMaterno", Value: aggregation.ArrayElemAt("$cliente.apellidoMaterno", 0)},
				{Key: "detallePago", Value: 1},
				{Key: "direccion", Value: aggregation.ArrayElemAt("$medidor.direccion", 0)},
				{Key: "codigoCliente", Value: aggregation.ArrayElemAt("$cliente.codigo", 0)},
			}},
		},
	}

	cursor, err := repo.collection.Aggregate(ctx, pipepine)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var data []bson.M = []bson.M{}
	err = cursor.All(ctx, &data)
	if err != nil {
		return nil, err
	}
	return &data[0], nil

}

func (repo *pagoRepository) ListarPagos(filter *dto.BuscardorPagoDto, ctx context.Context) (*map[string]interface{}, error) {
	var pipepine mongo.Pipeline = mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{

				{
					Key: "flag", Value: enum.FlagNuevo,
				},
			}},
		},
	}

	if filter.FechaInicio != "" && filter.FechaFin != "" {
		f1, f2, err := common.NormalizarRangoDeFechas(filter.FechaInicio, filter.FechaFin)

		if err != nil {
			return nil, err
		}

		pipepine = append(pipepine, bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "fecha", Value: bson.D{
					{Key: "$gte", Value: f1},
					{Key: "$lte", Value: f2},
				}},
			}},
		})
	}
	pipepine = append(pipepine, aggregation.Lookup("Cliente", "cliente", "_id", "cliente"))
	if filter.Ci != "" {
		pipepine = append(pipepine, aggregation.RegexMatch("cliente.0.ci", filter.Ci))
	}

	if filter.Nombre != "" {
		pipepine = append(pipepine, aggregation.RegexMatch("cliente.0.nombre", filter.Nombre))
	}

	if filter.ApellidoPaterno != "" {
		pipepine = append(pipepine, aggregation.RegexMatch("cliente.0.apellidoPaterno", filter.ApellidoPaterno))
	}

	if filter.ApellidoMaterno != "" {
		pipepine = append(pipepine, aggregation.RegexMatch("cliente.0.apellidoMaterno", filter.ApellidoMaterno))
	}

	if filter.CodigoCliente != "" {
		pipepine = append(pipepine, aggregation.RegexMatch("cliente.0.codigo", filter.CodigoCliente))
	}
	pipepine = append(pipepine, aggregation.Lookup("Medidor", "medidor", "_id", "medidor"))
	if filter.CodigoCliente != "" {
		pipepine = append(pipepine, aggregation.RegexMatch("medidor.0.numeroMedidor", filter.NumeroMedidor))
	}
	pipepine = append(pipepine,
		aggregation.Lookup("DetallePago", "_id", "pago", "detallePago"),
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "numeroPago", Value: 1},
				{Key: "total", Value: 1},
				{Key: "fecha", Value: 1},
				{Key: "numeroMedidor", Value: aggregation.ArrayElemAt("$medidor.numeroMedidor", 0)},
				{Key: "nombre", Value: aggregation.ArrayElemAt("$cliente.nombre", 0)},
				{Key: "apellidoPaterno", Value: aggregation.ArrayElemAt("$cliente.apellidoPaterno", 0)},
				{Key: "apellidoMaterno", Value: aggregation.ArrayElemAt("$cliente.apellidoMaterno", 0)},
				{Key: "detallePago", Value: 1},
				{Key: "codigoCliente", Value: aggregation.ArrayElemAt("$cliente.codigo", 0)},
				{Key: "ci", Value: aggregation.ArrayElemAt("$cliente.ci", 0)},
			}},
		},

		bson.D{
			{
				Key: "$facet", Value: bson.D{
					{Key: "data", Value: mongo.Pipeline{
						bson.D{{Key: "$skip", Value: aggregation.Skip(filter.Pagina, filter.Limite)}},
						bson.D{{Key: "$limit", Value: filter.Limite}},
					}},
					{Key: "countDocuments", Value: mongo.Pipeline{
						bson.D{{Key: "$count", Value: "countDocuments"}},
					}},
				},
			},
		})

	cursor, err := repo.collection.Aggregate(ctx, pipepine)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var resultado []aggregation.PaginacionResultado = []aggregation.PaginacionResultado{}
	err = cursor.All(ctx, &resultado)
	if err != nil {
		return nil, err
	}

	var total int = 0
	if len(resultado[0].CountDocuments) > 0 {
		total = int(resultado[0].CountDocuments[0].Count)
	}
	data := map[string]interface{}{
		"data":    resultado[0].Data,
		"paginas": aggregation.CalcularPaginas(total, filter.Limite),
		"total":   total,
		"pagina":  filter.Pagina,
		"limite":  filter.Limite,
	}
	return &data, nil
}

func (repo *pagoRepository) AnularPago(idPago *bson.ObjectID, cxt context.Context) (*model.Pago, error) {
	hoy := time.Now()
	pago, err := repo.BuscarPagoId(idPago, cxt)
	if err != nil {
		return nil, err
	}
	if pago.Fecha.Year() != hoy.Year() && pago.Fecha.Day() != hoy.Day() {
		return nil, fmt.Errorf("El pago solo se puede anular el mismo dia")
	}
	resultado, err := repo.collection.UpdateOne(cxt, bson.M{"_id": idPago}, bson.M{"$set": bson.M{"flag": enum.FlagAnulado, "FechaAnulacion": common.FechaHoraBolivia()}})
	if err != nil {
		return nil, err
	}
	if resultado.ModifiedCount > 0 {

		return &pago, nil
	}
	return nil, nil
}

func (repo *pagoRepository) BuscarPagosPorCaja(caja *bson.ObjectID, cxt context.Context) (*[]model.Pago, error) {
	cursor, err := repo.collection.Find(cxt, bson.M{"caja": caja, "flag": enum.FlagNuevo})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(cxt)
	var pagos []model.Pago = []model.Pago{}

	for cursor.Next(cxt) {
		var pago model.Pago = model.Pago{}
		err = cursor.Decode(&pago)
		if err != nil {
			return nil, err
		}
		pagos = append(pagos, pago)
	}

	err = cursor.Err()
	if err != nil {
		return nil, err
	}
	return &pagos, nil

}
