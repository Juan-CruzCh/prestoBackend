package common

import "math"

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func RedondearEfectivoBoliviano(valor float64) float64 {
	return math.Round(valor*10) / 10
}
