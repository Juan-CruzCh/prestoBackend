package repository

import (
	"context"
	"prestoBackend/src/app/enum"
	"prestoBackend/src/internal/pago/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type DetallePagoRepository interface {
	CrearDetalle(detalle *model.DetallePago, ctx context.Context) (*mongo.InsertOneResult, error)
	AnularDetallePago(pago *bson.ObjectID, ctx context.Context) (*[]model.DetallePago, error)
}

type detallePagoRepository struct {
	bd         *mongo.Database
	collection *mongo.Collection
}

func NewDetallePagoRepository(bd *mongo.Database) DetallePagoRepository {
	return &detallePagoRepository{
		bd:         bd,
		collection: bd.Collection("DetallePago"),
	}

}

func (repo *detallePagoRepository) CrearDetalle(detalle *model.DetallePago, ctx context.Context) (*mongo.InsertOneResult, error) {
	resultado, err := repo.collection.InsertOne(ctx, detalle)
	if err != nil {
		return nil, err
	}
	return resultado, nil
}

func (repo *detallePagoRepository) AnularDetallePago(pago *bson.ObjectID, ctx context.Context) (*[]model.DetallePago, error) {
	resultado, err := repo.collection.UpdateMany(ctx, bson.M{"pago": pago}, bson.M{"$set": bson.M{"flag": enum.FlagAnulado}})
	if err != nil {
		return nil, err
	}
	if resultado.ModifiedCount > 0 {
		cursor, err := repo.collection.Find(ctx, bson.M{"_id": pago, "flag": enum.FlagAnulado})
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)
		var detalles []model.DetallePago = []model.DetallePago{}

		for cursor.Next(ctx) {
			var detalle model.DetallePago = model.DetallePago{}
			err = cursor.Decode(&detalle)
			if err != nil {
				return nil, err
			}
			detalles = append(detalles, detalle)
		}
		err = cursor.Err()
		if err != nil {
			return nil, err
		}
		return &detalles, nil
	}
	return nil, nil

}
