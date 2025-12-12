package repository

import (
	"context"
	"errors"

	"prestoBackend/src/core/coreDto"
	"prestoBackend/src/core/enum"
	"prestoBackend/src/core/utils"
	"prestoBackend/src/module/cliente/dto"
	"prestoBackend/src/module/cliente/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type ClienteRepository interface {
	CrearCliente(cliente *model.Cliente, ctx context.Context) (*mongo.InsertOneResult, error)
	ListarClientes(filter dto.BucadorClienteDto, ctx context.Context) (*coreDto.ResultadoPaginado, error)
	ActualizarCliente()
	EliminarCliente()
	CantidadDocumentosCliente(ctx context.Context) (int, error)
	VerificarClienteCi(ci string, ctx context.Context) (int, error)
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

func (r *clienteRepository) CrearCliente(cliente *model.Cliente, ctx context.Context) (*mongo.InsertOneResult, error) {
	resultado, err := r.collection.InsertOne(ctx, cliente)
	if err != nil {
		return nil, err
	}
	return resultado, nil

}

func (r *clienteRepository) ListarClientes(filter dto.BucadorClienteDto, ctx context.Context) (*coreDto.ResultadoPaginado, error) {

	var filtro bson.D = bson.D{
		{Key: "flag", Value: enum.FlagNuevo},
	}
	if filter.Ci != "" {
		filtro = append(filtro, utils.Regex("ci", filter.Ci)...)
	}
	if filter.Nombre != "" {
		filtro = append(filtro, utils.Regex("nombre", filter.Nombre)...)
	}
	if filter.Codigo != "" {
		filtro = append(filtro, utils.Regex("codigo", filter.Codigo)...)
	}
	if filter.ApellidoPaterno != "" {
		filtro = append(filtro, utils.Regex("apellidoPaterno", filter.ApellidoPaterno)...)
	}

	if filter.ApellidoMaterno != "" {
		filtro = append(filtro, utils.Regex("apellidoMaterno", filter.ApellidoMaterno)...)
	}
	findOptions := options.Find()
	findOptions.SetSkip(int64(utils.Skip(filter.Pagina, filter.Limite)))
	findOptions.SetLimit(int64(filter.Limite))

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

	var paginas int = utils.CalcularPaginas(int(countDocuments), filter.Limite)
	var resultado coreDto.ResultadoPaginado = coreDto.ResultadoPaginado{
		Data:    &data,
		Total:   countDocuments,
		Paginas: paginas,
	}
	return &resultado, nil

}

func (r *clienteRepository) ActualizarCliente() {

}

func (r *clienteRepository) EliminarCliente() {

}

func (r *clienteRepository) CantidadDocumentosCliente(ctx context.Context) (int, error) {

	cantidad, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, errors.New("Se produjo un error en el conteo de clientes")
	}
	return int(cantidad), nil

}

func (r *clienteRepository) VerificarClienteCi(ci string, ctx context.Context) (int, error) {

	cantidad, err := r.collection.CountDocuments(ctx, bson.M{"ci": ci})
	if err != nil {
		return 0, errors.New("Se produjo un error al bsucar al cliete" + ci)
	}
	return int(cantidad), nil

}
