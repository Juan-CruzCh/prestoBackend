package utils

import (
	"math"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func Lookup(from, localField, foreignField, as string) bson.D {
	var pipelineMongo = bson.D{
		{
			Key: "$lookup",
			Value: bson.D{
				{Key: "from", Value: from},
				{Key: "localField", Value: localField},
				{Key: "foreignField", Value: foreignField},
				{Key: "as", Value: as},
			},
		},
	}
	return pipelineMongo
}

func Unwind(path string, preserveNullAndEmptyArrays bool) bson.D {
	return bson.D{
		{
			Key: "$unwind",
			Value: bson.D{
				{Key: "path", Value: path},
				{Key: "preserveNullAndEmptyArrays", Value: preserveNullAndEmptyArrays},
			},
		},
	}

}

func ArrayElemAt(arrayElemAtStage string, indice int) bson.D {
	var pipeline bson.D = bson.D{
		{Key: "$arrayElemAt", Value: bson.A{arrayElemAtStage, indice}},
	}
	return pipeline
}

func Regex(campo string, valor string) bson.D {
	return bson.D{{
		Key: campo,
		Value: bson.D{
			{Key: "$regex", Value: valor},
			{Key: "$options", Value: "i"},
		},
	}}
}
func RegexMatch(campo string, valor string) bson.D {
	return bson.D{{
		Key: "$match",
		Value: bson.D{{
			Key: campo,
			Value: bson.D{
				{Key: "$regex", Value: valor},
				{Key: "$options", Value: "i"},
			},
		}},
	}}
}

func Match(campo string, valor *bson.ObjectID) bson.D {
	return bson.D{{
		Key: "$match",
		Value: bson.D{{
			Key:   campo,
			Value: valor}},
	}}
}

func Skip(pagina, limite int) int {
	return (pagina - 1) * limite
}

func CalcularPaginas(total, limite int) int {
	if limite <= 0 {
		return 0
	}
	return int(math.Ceil(float64(total) / float64(limite)))
}

func Sort(campo string) bson.D {
	return bson.D{
		{Key: "$sort", Value: bson.D{
			{Key: campo, Value: -1},
		}},
	}
}

type PaginacionResultado struct {
	Data           []bson.M `bson:"data"`
	CountDocuments []struct {
		Count int64 `bson:"countDocuments"`
	} `bson:"countDocuments"`
}
