package repository

import "go.mongodb.org/mongo-driver/v2/mongo"

type MedidorRepository interface {
	ListarMedidor()
	EliminarMedidor()
	ActualizarMedidor()
	crearMedidor()
}

type medidorRepository struct {
	db         *mongo.Database
	Collection *mongo.Collection
}

func NewMedidorRespository(db *mongo.Database) MedidorRepository {
	return &medidorRepository{
		db:         db,
		Collection: db.Collection("Lectura"),
	}
}

func (r *medidorRepository) ListarMedidor() {

}

func (r *medidorRepository) EliminarMedidor() {

}

func (r *medidorRepository) ActualizarMedidor() {

}

func (r *medidorRepository) crearMedidor() {

}
