package repository

import (
	"context"
	"errors"
	"strconv"

	"prestoBackend/src/app/common"
	"prestoBackend/src/app/database/aggregation"
	"prestoBackend/src/app/enum"
	"prestoBackend/src/internal/cliente/dto"
	"prestoBackend/src/internal/cliente/model"

	appDto "prestoBackend/src/app/dto"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type ClienteRepository interface {
	CrearCliente(cliente *model.Cliente, ctx context.Context) (*model.Cliente, error)
	ListarClientes(filter dto.BucadorClienteDto, ctx context.Context) (*appDto.ResultadoPaginado, error)
	ActualizarCliente(cliente *model.Cliente, ID *bson.ObjectID, ctx context.Context) (*mongo.UpdateResult, error)
	EliminarCliente(ID *bson.ObjectID, ctx context.Context) (*mongo.UpdateResult, error)
}

type clienteRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewClienteRepository(db *mongo.Database) ClienteRepository {
	return &clienteRepository{
		db:         db,
		collection: db.Collection("Cliente"),
	}
}

func (r *clienteRepository) CrearCliente(cliente *model.Cliente, ctx context.Context) (*model.Cliente, error) {

	cantidadCliente, err := r.collection.CountDocuments(ctx, bson.M{"ci": cliente.Ci, "flag": enum.FlagNuevo})
	if err != nil {
		return nil, errors.New("Se produjo un error al bsucar al cliete" + cliente.Ci)
	}
	if cantidadCliente > 0 {
		return nil, errors.New("El cliente ya se encuetra registrado")
	}
	cantidad, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, errors.New("Se produjo un error en el conteo de clientes")
	}
	cliente.Codigo = "C-" + strconv.Itoa(int(cantidad))
	cliente.Flag = enum.FlagNuevo
	cliente.Fecha = common.FechaHoraBolivia()
	resultado, err := r.collection.InsertOne(ctx, cliente)
	if err != nil {
		return nil, err
	}
	cliente.ID = resultado.InsertedID.(bson.ObjectID)
	return cliente, nil

}

func (r *clienteRepository) ListarClientes(filter dto.BucadorClienteDto, ctx context.Context) (*appDto.ResultadoPaginado, error) {

	var filtro bson.D = bson.D{
		{Key: "flag", Value: enum.FlagNuevo},
	}
	if filter.Ci != "" {
		filtro = append(filtro, aggregation.Regex("ci", filter.Ci)...)
	}
	if filter.Nombre != "" {
		filtro = append(filtro, aggregation.Regex("nombre", filter.Nombre)...)
	}
	if filter.Codigo != "" {
		filtro = append(filtro, aggregation.Regex("codigo", filter.Codigo)...)
	}
	if filter.ApellidoPaterno != "" {
		filtro = append(filtro, aggregation.Regex("apellidoPaterno", filter.ApellidoPaterno)...)
	}

	if filter.ApellidoMaterno != "" {
		filtro = append(filtro, aggregation.Regex("apellidoMaterno", filter.ApellidoMaterno)...)
	}
	if filter.Celular != "" {
		filtro = append(filtro, aggregation.Regex("celular", filter.Celular)...)
	}
	findOptions := options.Find()
	findOptions.SetSkip(int64(aggregation.Skip(filter.Pagina, filter.Limite)))
	findOptions.SetLimit(int64(filter.Limite))
	findOptions.SetSort(bson.M{"fecha": -1})
	cursor, err := r.collection.Find(ctx, filtro, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var data []bson.M = []bson.M{}
	err = cursor.All(ctx, &data)
	if err != nil {
		return nil, err
	}

	countDocuments, err := r.collection.CountDocuments(ctx, filtro)
	if err != nil {
		return nil, err
	}

	var paginas int = aggregation.CalcularPaginas(int(countDocuments), filter.Limite)
	var resultado appDto.ResultadoPaginado = appDto.ResultadoPaginado{
		Data:    &data,
		Total:   countDocuments,
		Paginas: paginas,
	}
	return &resultado, nil

}

func (r *clienteRepository) ActualizarCliente(cliente *model.Cliente, ID *bson.ObjectID, ctx context.Context) (*mongo.UpdateResult, error) {
	var filter bson.D = bson.D{
		{Key: "ci", Value: cliente.Ci},
		{Key: "_id", Value: bson.D{
			{Key: "$ne", Value: ID},
		}},
		{Key: "flag", Value: enum.FlagNuevo},
	}
	cantidadCliente, err := r.collection.CountDocuments(ctx, filter)
	if err != nil {

		return nil, errors.New("Se produjo un error al buscar al cliente" + cliente.Ci)
	}
	if cantidadCliente > 0 {
		return nil, errors.New("El cliente ya se encuetra registrado")
	}

	var actualizar bson.D = bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "ci", Value: cliente.Ci},
			{Key: "nombre", Value: cliente.Nombre},
			{Key: "apellidoMaterno", Value: cliente.ApellidoMaterno},
			{Key: "apellidoPaterno", Value: cliente.ApellidoPaterno},
			{Key: "celular", Value: cliente.Celular},
		}},
	}
	restultado, err := r.collection.UpdateOne(ctx, bson.M{"flag": enum.FlagNuevo, "_id": ID}, actualizar)
	if err != nil {
		return nil, err
	}
	return restultado, nil
}

func (r *clienteRepository) EliminarCliente(ID *bson.ObjectID, ctx context.Context) (*mongo.UpdateResult, error) {
	var flagEliminado bson.D = bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "flag", Value: enum.FlagEliminado},
		}},
	}
	restultado, err := r.collection.UpdateOne(ctx, bson.M{"flag": enum.FlagNuevo, "_id": ID}, flagEliminado)
	if err != nil {
		return nil, err
	}

	return restultado, nil
}
