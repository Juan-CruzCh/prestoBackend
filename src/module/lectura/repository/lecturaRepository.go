package repository

import "go.mongodb.org/mongo-driver/v2/mongo"

type LecturaRepository interface {
	CrearLectura()
	ListarLectura()
	ActualizarLectura()
	EliminarLectuta()
}

type lecturaRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewLecturaRepository(db *mongo.Database) LecturaRepository {
	return &lecturaRepository{
		db:         db,
		collection: db.Collection("Lectura"),
	}
}

func (r *lecturaRepository) CrearLectura() {

}

func (r *lecturaRepository) ListarLectura() {

}
func (r *lecturaRepository) ActualizarLectura() {

}
func (r *lecturaRepository) EliminarLectuta() {

}
