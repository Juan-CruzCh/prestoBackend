package repository

import (
	"context"
	"errors"
	"prestoBackend/src/module/cliente/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ClienteRepository interface {
	CrearCliente(cliente *model.Cliente, ctx context.Context) (*mongo.InsertOneResult, error)
	ListarClientes()
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

func (r *clienteRepository) ListarClientes() {

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
