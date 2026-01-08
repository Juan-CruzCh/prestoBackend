package repository

import (
	"context"
	"fmt"
	"prestoBackend/src/core/enum"
	"prestoBackend/src/module/tarifa/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type TarifaRepository interface {
	CrearTarifa(tarifa *model.Tarifa, ctx context.Context) (*mongo.InsertOneResult, error)
	VerificarTarifa(nombre string, ctx context.Context) (int, error)
	ListarTarifas(ctx context.Context) ([]model.Tarifa, error)
	EliminarTarifa(tarifa *bson.ObjectID, ctx context.Context) (*mongo.UpdateResult, error)
}

type tarifaRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewTarifaRepository(db *mongo.Database) TarifaRepository {
	return &tarifaRepository{
		db:         db,
		collection: db.Collection("Tarifa"),
	}

}
func (r *tarifaRepository) CrearTarifa(tarifa *model.Tarifa, ctx context.Context) (*mongo.InsertOneResult, error) {
	resultado, err := r.collection.InsertOne(ctx, tarifa)
	if err != nil {
		return nil, err
	}
	return resultado, nil
}

func (r *tarifaRepository) VerificarTarifa(nombre string, ctx context.Context) (int, error) {
	cantidad, err := r.collection.CountDocuments(ctx, bson.M{"nombre": nombre})
	if err != nil {
		return 0, err
	}
	return int(cantidad), nil
}

func (r *tarifaRepository) ListarTarifas(ctx context.Context) ([]model.Tarifa, error) {

	cursor, err := r.collection.Find(ctx, bson.M{"flag": enum.FlagNuevo})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var tarifas []model.Tarifa = []model.Tarifa{}
	err = cursor.All(ctx, &tarifas)
	if err != nil {
		return nil, err
	}

	return tarifas, nil
}

func (repository *tarifaRepository) EliminarTarifa(tarifa *bson.ObjectID, ctx context.Context) (*mongo.UpdateResult, error) {
	var flagEliminado bson.D = bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "flag", Value: enum.FlagEliminado},
		}},
	}
	resultado, err := repository.collection.UpdateOne(ctx, bson.M{"flag": enum.FlagNuevo, "_id": tarifa}, flagEliminado)
	if err != nil {
		return nil, err
	}
	if resultado.MatchedCount == 0 {
		return nil, fmt.Errorf("La tarifa no existe")
	}
	return resultado, nil

}
