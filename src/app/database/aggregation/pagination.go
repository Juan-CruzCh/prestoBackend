package aggregation

import (
	"math"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func Skip(pagina, limite int) int {
	return (pagina - 1) * limite
}

func CalcularPaginas(total, limite int) int {
	if limite <= 0 {
		return 0
	}
	return int(math.Ceil(float64(total) / float64(limite)))
}

type PaginacionResultado struct {
	Data           []bson.M `bson:"data"`
	CountDocuments []struct {
		Count int64 `bson:"countDocuments"`
	} `bson:"countDocuments"`
}
