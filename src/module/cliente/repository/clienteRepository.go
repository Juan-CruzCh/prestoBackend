package repository

import "go.mongodb.org/mongo-driver/v2/mongo"

type ClienteRepository interface {
	CrearCliente()
	ListarClientes()
	ActualizarCliente()
	EliminarCliente()
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

func (r *clienteRepository) CrearCliente() {

}

func (r *clienteRepository) ListarClientes() {

}

func (r *clienteRepository) ActualizarCliente() {

}

func (r *clienteRepository) EliminarCliente() {

}
