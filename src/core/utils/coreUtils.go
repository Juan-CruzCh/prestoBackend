package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func ValidadIdMongo(id string) (*bson.ObjectID, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("ID de mongo invalido")
	}
	return &objectID, nil
}

func ValidarMongoIdArray(id []string) (IDS []bson.ObjectID, err error) {
	var ids []bson.ObjectID
	for _, v := range id {
		objID, err := ValidadIdMongo(v)
		if err != nil {
			return nil, err
		}
		ids = append(ids, *objID)
	}
	return ids, nil
}

func FechaHoraBolivia() time.Time {
	fecha := time.Now()
	return fecha.Add(-4 * time.Hour)

}

func PrintLnCustomArray(valor *[]bson.M) {
	jsonData, err := json.MarshalIndent(valor, "", "  ")
	if err != nil {
		fmt.Println("Ocrrui un error")
	}
	fmt.Println(string(jsonData))
}

func PrintLnCustom(valor *bson.M) {
	jsonData, err := json.MarshalIndent(valor, "", "  ")
	if err != nil {
		fmt.Println("Ocrrui un error")
	}
	fmt.Println(string(jsonData))
}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
