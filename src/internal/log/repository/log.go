package repository

import (
	"context"
	"prestoBackend/src/app/common"
	"prestoBackend/src/app/enum"
	"prestoBackend/src/internal/log/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Log interface {
	CrearLog(log *model.Log, ctx context.Context) error
	ListarLog(ctx context.Context) (*[]model.Log, error)
}

type log struct {
	collection *mongo.Collection
}

func NewLogRepository(db *mongo.Database) *log {
	collection := db.Collection("Log")
	return &log{collection: collection}
}

func (r *log) CrearLog(log *model.Log, ctx context.Context) error {
	log.Fecha = common.FechaHoraBolivia()
	log.Flag = enum.FlagNuevo
	_, err := r.collection.InsertOne(ctx, log)
	if err != nil {
		return err
	}
	return nil
}

func (r *log) ListarLog(ctx context.Context) (*[]model.Log, error) {
	cursor, err := r.collection.Find(ctx, bson.M{"flag": enum.FlagNuevo})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var resultado []model.Log = []model.Log{}
	for cursor.Next(ctx) {
		var log model.Log
		err = cursor.Decode(&log)
		if err != nil {
			return nil, err
		}
		resultado = append(resultado, log)
	}
	err = cursor.Err()
	if err != nil {
		return nil, err
	}

	return &resultado, nil
}
