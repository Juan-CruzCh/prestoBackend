package repository

import (
	"context"
	"fmt"
	"prestoBackend/src/core/enum"
	"prestoBackend/src/module/usuario/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UsuarioRepository interface {
	CrearUsuario(usuario *model.Usuario, ctx context.Context) (*mongo.InsertOneResult, error)
	ListarUsuario(ctx context.Context) (*[]model.Usuario, error)
	EliminarUsuario(usuario *bson.ObjectID, ctx context.Context) (*mongo.UpdateResult, error)
}

type usuarioRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewUsuarioRepository(db *mongo.Database) UsuarioRepository {
	return &usuarioRepository{db: db, collection: db.Collection("Usuario")}
}

func (repo *usuarioRepository) CrearUsuario(usuario *model.Usuario, ctx context.Context) (*mongo.InsertOneResult, error) {
	cantidad, err := repo.collection.CountDocuments(ctx, bson.M{"flag": enum.FlagNuevo, "usuario": usuario.Usuario})
	if err != nil {
		return nil, err
	}
	if cantidad > 0 {
		return nil, fmt.Errorf("El usuario ya existe")
	}
	resultado, err := repo.collection.InsertOne(ctx, usuario)
	if err != nil {
		return nil, err
	}
	return resultado, nil
}

func (repo *usuarioRepository) ListarUsuario(ctx context.Context) (*[]model.Usuario, error) {
	cursor, err := repo.collection.Find(ctx, bson.M{"flag": enum.FlagNuevo})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var data []model.Usuario = []model.Usuario{}
	err = cursor.All(ctx, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (repo *usuarioRepository) BuscarUsuarioPorUsuario(usuario string, ctx context.Context) (*model.Usuario, error) {
	var data model.Usuario
	err := repo.collection.FindOne(ctx, bson.M{"flag": enum.FlagNuevo, "usuario": usuario}).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
func (repository *usuarioRepository) EliminarUsuario(usuario *bson.ObjectID, ctx context.Context) (*mongo.UpdateResult, error) {
	var flagEliminado bson.D = bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "flag", Value: enum.FlagEliminado},
		}},
	}
	resultado, err := repository.collection.UpdateOne(ctx, bson.M{"flag": enum.FlagNuevo, "_id": usuario}, flagEliminado)
	if err != nil {
		return nil, err
	}
	if resultado.MatchedCount == 0 {
		return nil, fmt.Errorf("El usuario no existe")
	}
	return resultado, nil
}
