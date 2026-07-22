package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type usuarioAutenticado struct {
	ID  bson.ObjectID
	Rol string
}

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

func ObtenerUsuarioRequest(w http.ResponseWriter, r *http.Request) (*usuarioAutenticado, error) {
	usuarioContext := r.Context().Value("usuario")
	usuario, ok := usuarioContext.(map[string]string)
	if !ok {
		return nil, fmt.Errorf("Usuario no encontrado en contexto")
	}
	usuarioID, err := ValidadIdMongo(usuario["id"])
	if err != nil {

		return nil, err
	}
	var resultado usuarioAutenticado = usuarioAutenticado{
		ID:  *usuarioID,
		Rol: usuario["rol"],
	}
	return &resultado, nil
}
