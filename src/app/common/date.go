package common

import (
	"errors"
	"time"
)

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

func FechaHoraBolivia() time.Time {
	fecha := time.Now()
	return fecha.Add(-4 * time.Hour)

}
