package coreDto

import "go.mongodb.org/mongo-driver/v2/bson"

type ResultadoPaginado struct {
	Data    *[]bson.M `json:"data"`
	Total   int64     `json:"total"`
	Paginas int       `json:"paginas"`
}
