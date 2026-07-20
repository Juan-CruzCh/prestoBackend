package common

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func ResponseJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
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
