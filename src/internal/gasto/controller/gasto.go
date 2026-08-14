package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"prestoBackend/src/app/common"
	"prestoBackend/src/internal/gasto/dto"
	"prestoBackend/src/internal/gasto/service"
	"time"

	"github.com/go-playground/validator/v10"
)

type Gasto struct {
	gastoService *service.Gasto
	Validate     *validator.Validate
}

func NewGastoController(gastoService *service.Gasto, Validate *validator.Validate) *Gasto {
	return &Gasto{
		gastoService: gastoService,
	}
}
func (c *Gasto) CrearGasto(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	usuario, err := common.ObtenerUsuarioRequest(w, r)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	var body dto.GastoDto
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	err = c.Validate.Struct(&body)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	err = c.gastoService.CrearGasto(&body, &usuario.ID, ctx)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	common.ResponseJSON(w, http.StatusCreated, map[string]string{"mensaje": "Gasto creado"})
}

func (c *Gasto) ListarGasto(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	resultado, err := c.gastoService.ListarGasto(ctx)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	common.ResponseJSON(w, http.StatusOK, resultado)

}

func (c *Gasto) EliminarGasto(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	usuario, err := common.ObtenerUsuarioRequest(w, r)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}

	var idPago string = r.PathValue("id")
	ID, err := common.ValidadIdMongo(idPago)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	err = c.gastoService.EliminarGasto(ID, &usuario.ID, ctx)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	common.ResponseJSON(w, http.StatusOK, map[string]string{"mensaje": "Eliminado"})

}
