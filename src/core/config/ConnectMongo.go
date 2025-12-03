package config

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func ConnectMongo(url string, nombreBd string) (*mongo.Database, *mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	cliente, err := mongo.Connect(options.Client().ApplyURI(url))

	if err != nil {
		return nil, nil, err
	}
	err = cliente.Ping(ctx, nil)

	fmt.Println("ping", err)
	if err != nil {
		return nil, nil, err
	}

	fmt.Println("Se conecto a la base de datos")
	return cliente.Database(nombreBd), cliente, nil

}
