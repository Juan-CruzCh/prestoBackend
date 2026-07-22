package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"prestoBackend/src/app/common"
	"prestoBackend/src/internal/cajaChica/dto"
	"prestoBackend/src/internal/cajaChica/service"
	"time"

	"github.com/go-playground/validator/v10"
)

type CajaChica struct {
	cajaChicaService *service.CajaChica
	Validate         *validator.Validate
}

func NewCajaChicaController(cajaChicaService *service.CajaChica, Validate *validator.Validate) *CajaChica {
	return &CajaChica{
		cajaChicaService: cajaChicaService,
		Validate:         Validate,
	}
}
func (c *CajaChica) CrearCajaChica(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var body dto.CajaChicaDto
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	err = c.Validate.Struct(&body)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}

	usuario, err := common.ObtenerUsuarioRequest(w, r)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	err = c.cajaChicaService.CrearCajaChica(&body, usuario.ID, ctx)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": "Caja abierta"})
}

func (c *CajaChica) ListarCajaChica(w http.ResponseWriter, r *http.Request) {
}

func (c *CajaChica) ActualizarCajaChica(w http.ResponseWriter, r *http.Request) {
}

func (c *CajaChica) EliminarCajaChica(w http.ResponseWriter, r *http.Request) {
}
