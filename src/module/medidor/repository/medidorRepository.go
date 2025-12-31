package repository

import (
	"context"
	"errors"
	"fmt"
	"prestoBackend/src/core/coreDto"
	"prestoBackend/src/core/enum"
	"prestoBackend/src/core/utils"
	"prestoBackend/src/module/medidor/dto"
	"prestoBackend/src/module/medidor/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type MedidorRepository interface {
	ActualizarMedidor()
	CrearMedidor(medidor *model.Medidor, ctx context.Context) (*mongo.InsertOneResult, error)
	CantidadMedidor(ctx context.Context) (int, error)
	ObtenerMedidor(medidor *bson.ObjectID, ctx context.Context) (*model.Medidor, error)
	ActualizaLecturasPendientesMedidor(cantidad int, medidor *bson.ObjectID, ctx context.Context) error
	ListarMedidorCliente(filter *dto.BuscadorMedidorClienteDto, ctx context.Context) (*coreDto.ResultadoPaginado, error)
	BuscarMedidorPorNumeroMedidor(numeroMedidor string, ctx context.Context) ([]dto.MedidorClienteProject, error)
	BuscarMedidorCliente(cliente *bson.ObjectID, ctx context.Context) ([]model.Medidor, error)
	ObtenerMedidorConCliente(medidor *bson.ObjectID, ctx context.Context) (*[]bson.M, error)
	EliminarMedidor(medidor *bson.ObjectID, ctx context.Context) (*mongo.UpdateResult, error)
	EliminarMedidoresCliente(cliente *bson.ObjectID, ctx context.Context) (*mongo.UpdateResult, error)
}

type medidorRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewMedidorRespository(db *mongo.Database) MedidorRepository {
	return &medidorRepository{
		db:         db,
		collection: db.Collection("Medidor"),
	}
}

func (r *medidorRepository) ActualizarMedidor() {

}

func (r *medidorRepository) CrearMedidor(medidor *model.Medidor, ctx context.Context) (*mongo.InsertOneResult, error) {

	cantidad, err := r.collection.CountDocuments(ctx, bson.M{"numeroMedidor": medidor.NumeroMedidor})
	if err != nil {
		return nil, err
	}
	if cantidad > 0 {
		return nil, errors.New("el medidor ya se encuentra registrado")
	}
	resultado, err := r.collection.InsertOne(ctx, medidor)
	if err != nil {
		return nil, err
	}
	return resultado, nil

}

func (r *medidorRepository) CantidadMedidor(ctx context.Context) (int, error) {

	cantidad, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}

	return int(cantidad), nil

}

func (r *medidorRepository) ObtenerMedidor(medidor *bson.ObjectID, ctx context.Context) (*model.Medidor, error) {
	var data model.Medidor
	err := r.collection.FindOne(ctx, bson.M{"_id": medidor, "flag": enum.FlagNuevo}).Decode(&data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &data, nil

}

func (r *medidorRepository) ActualizaLecturasPendientesMedidor(cantidad int, medidor *bson.ObjectID, ctx context.Context) error {

	var filtro bson.M = bson.M{"_id": *medidor}

	var update bson.D = bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "lecturasPendientes", Value: cantidad},
		}},
	}
	_, err := r.collection.UpdateOne(ctx, filtro, update)

	return err

}

func (r *medidorRepository) ListarMedidorCliente(filter *dto.BuscadorMedidorClienteDto, ctx context.Context) (*coreDto.ResultadoPaginado, error) {

	var pipeline mongo.Pipeline = mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "flag", Value: enum.FlagNuevo},
			}},
		},
	}
	if filter.NumeroMedidor != "" {
		pipeline = append(pipeline, utils.RegexMatch("numeroMedidor", filter.NumeroMedidor))
	}

	if filter.EstadoMedidor == "moroso" {
		pipeline = append(pipeline, bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "lecturasPendientes", Value: bson.D{
					{Key: "$gt", Value: 3},
				}},
			}},
		})
	}

	if filter.Estado != "" {
		pipeline = append(pipeline, bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "estado", Value: filter.Estado},
			}},
		})
	}
	if filter.Direccion != "" {
		pipeline = append(pipeline, utils.RegexMatch("direccion", filter.Direccion))
	}
	pipeline = append(pipeline,
		utils.Lookup("Cliente", "cliente", "_id", "cliente"),
		utils.Unwind("$cliente", false),
	)

	if filter.Codigo != "" {
		pipeline = append(pipeline, utils.RegexMatch("cliente.codigo", filter.Codigo))
	}
	if filter.Nombre != "" {
		pipeline = append(pipeline, utils.RegexMatch("cliente.nombre", filter.Nombre))
	}

	if filter.ApellidoPaterno != "" {
		pipeline = append(pipeline, utils.RegexMatch("cliente.apellidoPaterno", filter.ApellidoPaterno))
	}
	if filter.ApellidoMaterno != "" {
		pipeline = append(pipeline, utils.RegexMatch("cliente.apellidoMaterno", filter.ApellidoMaterno))
	}

	pipeline = append(pipeline,
		utils.Lookup("Tarifa", "tarifa", "_id", "tarifa"),
		utils.Unwind("$tarifa", false),
	)
	if filter.Tarifa != "" {
		ID, err := utils.ValidadIdMongo(filter.Tarifa)
		if err != nil {
			return nil, err
		}
		fmt.Println(ID)
		pipeline = append(pipeline, utils.Match("tarifa._id", ID))
	}

	pipeline = append(pipeline,
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 1},
				{Key: "numeroMedidor", Value: 1},
				{Key: "lecturasPendientes", Value: 1},
				{Key: "estado", Value: 1},
				{Key: "direccion", Value: 1},
				{Key: "nombre", Value: "$cliente.nombre"},
				{Key: "apellidoPaterno", Value: "$cliente.apellidoPaterno"},
				{Key: "apellidoMaterno", Value: "$cliente.apellidoMaterno"},
				{Key: "codigo", Value: "$cliente.codigo"},
				{Key: "ci", Value: "$cliente.ci"},
				{Key: "tarifa", Value: "$tarifa.nombre"},
			}},
		},
	)

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {

		return nil, err
	}
	defer cursor.Close(ctx)

	countDocuments, err := r.collection.CountDocuments(ctx, bson.M{"flag": enum.FlagNuevo})
	if err != nil {

		return nil, err
	}
	var data []bson.M = []bson.M{}

	err = cursor.All(ctx, &data)
	if err != nil {

		return nil, err
	}

	var paginas int = utils.CalcularPaginas(int(countDocuments), filter.Limite)
	var resultado coreDto.ResultadoPaginado = coreDto.ResultadoPaginado{
		Data:    &data,
		Total:   countDocuments,
		Paginas: paginas,
	}
	return &resultado, nil
}

func (r *medidorRepository) ListarMedidorClienteMorosos(ctx context.Context) (*[]bson.M, error) {
	var pipeline mongo.Pipeline = mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "flag", Value: enum.FlagNuevo},
				{Key: "LecturasPendientes", Value: bson.D{
					{Key: "$gt", Value: 3},
				}},
			}},
		},
		utils.Lookup("Cliente", "cliente", "_id", "cliente"),
		utils.Unwind("$cliente", false),
		utils.Lookup("Tarifa", "tarifa", "_id", "tarifa"),

		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 1},
				{Key: "numeroMedidor", Value: 1},
				{Key: "estado", Value: 1},
				{Key: "direccion", Value: 1},
				{Key: "nombre", Value: "$cliente.nombre"},
				{Key: "apellidoPaterno", Value: "$cliente.apellidoPaterno"},
				{Key: "apellidoMaterno", Value: "$cliente.apellidoMaterno"},
				{Key: "codigo", Value: "$cliente.codigo"},
				{Key: "tarifa", Value: utils.ArrayElemAt("$tarifa.nombre", 0)},
			}},
		},
	}
	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {

		return nil, err
	}
	defer cursor.Close(ctx)

	var data []bson.M = []bson.M{}

	err = cursor.All(ctx, &data)
	if err != nil {

		return nil, err
	}

	return &data, nil
}

func (r *medidorRepository) BuscarMedidorPorNumeroMedidor(numeroMedidor string, ctx context.Context) ([]dto.MedidorClienteProject, error) {

	var pipeline mongo.Pipeline = mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{
				{Key: "numeroMedidor", Value: numeroMedidor},
				{Key: "flag", Value: enum.FlagNuevo},
			}},
		},
		utils.Lookup("Cliente", "cliente", "_id", "cliente"),
		bson.D{
			{
				Key: "$project", Value: bson.D{
					{Key: "nombre", Value: utils.ArrayElemAt("$cliente.nombre", 0)},
					{Key: "apellidoPaterno", Value: utils.ArrayElemAt("$cliente.apellidoPaterno", 0)},
					{Key: "apellidoMaterno", Value: utils.ArrayElemAt("$cliente.apellidoMaterno", 0)},
					{Key: "numeroMedidor", Value: 1},
					{Key: "estado", Value: 1},
					{Key: "_id", Value: 1},
				},
			},
		},
	}
	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var resultado []dto.MedidorClienteProject

	err = cursor.All(ctx, &resultado)
	if err != nil {
		return nil, err
	}
	return resultado, nil
}

func (r *medidorRepository) BuscarMedidorCliente(cliente *bson.ObjectID, ctx context.Context) ([]model.Medidor, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"cliente": cliente, "flag": enum.FlagNuevo})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var medidor []model.Medidor = []model.Medidor{}
	err = cursor.All(ctx, &medidor)
	if err != nil {
		return nil, err
	}
	return medidor, nil

}

func (r *medidorRepository) ObtenerMedidorConCliente(medidor *bson.ObjectID, ctx context.Context) (*[]bson.M, error) {

	var pipeline mongo.Pipeline = mongo.Pipeline{
		bson.D{
			{Key: "$match", Value: bson.D{
				{
					Key: "_id", Value: medidor,
				},
				{
					Key: "flag", Value: enum.FlagNuevo,
				},
			}},
		},
		utils.Lookup("Cliente", "cliente", "_id", "cliente"),
		utils.Lookup("Tarifa", "tarifa", "_id", "tarifa"),
		bson.D{
			{
				Key: "$project", Value: bson.D{
					{Key: "nombre", Value: utils.ArrayElemAt("$cliente.nombre", 0)},
					{Key: "apellidoPaterno", Value: utils.ArrayElemAt("$cliente.apellidoPaterno", 0)},
					{Key: "apellidoMaterno", Value: utils.ArrayElemAt("$cliente.apellidoMaterno", 0)},
					{Key: "numeroMedidor", Value: 1},
					{Key: "direccion", Value: 1},
					{Key: "codigoCliente", Value: utils.ArrayElemAt("$cliente.codigo", 0)},
					{Key: "tarifa", Value: utils.ArrayElemAt("$tarifa.nombre", 0)},
					{Key: "_id", Value: 1},
				},
			},
		},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var data []bson.M = []bson.M{}
	err = cursor.All(ctx, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil

}
func (r *medidorRepository) EliminarMedidor(medidor *bson.ObjectID, ctx context.Context) (*mongo.UpdateResult, error) {
	var flagEliminado bson.D = bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "flag", Value: enum.FlagEliminado},
		}},
	}
	resultado, err := r.collection.UpdateOne(ctx, bson.M{"_id": medidor}, flagEliminado)
	if err != nil {
		return nil, err
	}
	return resultado, nil

}
func (r *medidorRepository) EliminarMedidoresCliente(cliente *bson.ObjectID, ctx context.Context) (*mongo.UpdateResult, error) {
	var flagEliminado bson.D = bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "flag", Value: enum.FlagEliminado},
		}},
	}
	resultado, err := r.collection.UpdateMany(ctx, bson.M{"cliente": cliente}, flagEliminado)
	if err != nil {
		return nil, err
	}
	return resultado, nil

}
