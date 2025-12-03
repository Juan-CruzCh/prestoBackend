package repository

import "go.mongodb.org/mongo-driver/v2/mongo"

type UsuarioRepository interface {
	CrearUsuario()
}

type usuarioRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewUsuarioRepository(db *mongo.Database) UsuarioRepository {
	return &usuarioRepository{db: db, collection: db.Collection("Usuario")}
}

func (repo *usuarioRepository) CrearUsuario() {

}
