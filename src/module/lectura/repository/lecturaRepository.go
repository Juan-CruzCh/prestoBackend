package repository

import (
	"context"
	"prestoBackend/src/module/lectura/dto"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type LecturaRepository interface {
	CrearLectura(lecturaDto *dto.LecturaDto, ctx context.Context)
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

func (r *lecturaRepository) CrearLectura(lecturaDto *dto.LecturaDto, ctx context.Context) {

}

func (r *lecturaRepository) ListarLectura() {

}
func (r *lecturaRepository) ActualizarLectura() {

}
func (r *lecturaRepository) EliminarLectuta() {

}
