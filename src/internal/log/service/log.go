package service

import (
	"context"
	"prestoBackend/src/internal/log/model"
	"prestoBackend/src/internal/log/repository"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Log struct {
	logRepository repository.Log
	cliente       *mongo.Client
}

func NewLogService(logRepository repository.Log, cliente *mongo.Client) *Log {
	return &Log{
		logRepository: logRepository,
		cliente:       cliente,
	}
}

func (s *Log) CrearLog(ctx context.Context) error {
	return nil
}

func (s *Log) ListarLog(ctx context.Context) (*[]model.Log, error) {
	data, err := s.logRepository.ListarLog(ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}
