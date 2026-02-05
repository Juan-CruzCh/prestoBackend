package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
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

func PaginadorHTTP(r *http.Request) (int, int, error) {
	query := r.URL.Query()
	paginaStr := query.Get("pagina")
	if paginaStr == "" {
		paginaStr = "1"
	}
	limiteStr := query.Get("limite")
	if limiteStr == "" {
		limiteStr = "20"
	}

	pagina, err := strconv.Atoi(paginaStr)
	if err != nil {
		return 0, 0, errors.New("Ingrese el numero pagina")
	}

	limite, err := strconv.Atoi(limiteStr)
	if err != nil {
		return 0, 0, errors.New("Ingrese el numero limite")
	}

	return pagina, limite, nil
}

func NormalizarRangoDeFechas(fechaInicio string, fechaFin string) (f1 time.Time, f2 time.Time, err error) {
	const layout = "2006-01-02"

	parsedInicio, err1 := time.Parse(layout, fechaInicio)
	if err1 != nil {

		return f1, f2, errors.New("error fecha inicio" + err1.Error())
	}

	parsedFin, err2 := time.Parse(layout, fechaFin)
	if err2 != nil {
		return f1, f2, errors.New("error fecha fil" + err2.Error())

	}

	f1 = time.Date(parsedInicio.Year(), parsedInicio.Month(), parsedInicio.Day(), 0, 0, 0, 0, time.UTC)
	f2 = time.Date(parsedFin.Year(), parsedFin.Month(), parsedFin.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), time.UTC)

	return f1, f2, nil
}

func RedondearEfectivoBoliviano(valor float64) float64 {
	return math.Round(valor*10) / 10
}

func ErrorJson(err error) []map[string]string {
	errores := make([]map[string]string, 0)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {

		for _, e := range validationErrors {
			errorMsg := e.Tag()
			if e.Param() != "" {
				errorMsg += "=" + e.Param()
			}

			errores = append(errores, map[string]string{
				"field": e.Field(),
				"error": errorMsg,
			})
		}

		return errores
	}
	return errores
}

func ResponseJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}
