package common

import (
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

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
