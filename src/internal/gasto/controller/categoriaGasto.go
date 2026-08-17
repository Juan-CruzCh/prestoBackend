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

type CategoriaGasto struct {
	categoriaGastoService *service.CategoriaGasto
	Validate              *validator.Validate
}

func NewCategoriaGastoController(categoriaGastoService *service.CategoriaGasto, Validate *validator.Validate) *CategoriaGasto {
	return &CategoriaGasto{
		categoriaGastoService: categoriaGastoService,
		Validate:              Validate,
	}
}
func (c *CategoriaGasto) CrearCategoriaGasto(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var body dto.CategoriaGastoDto
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

	err = c.categoriaGastoService.CrearCategoriaGasto(&body, ctx)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	common.ResponseJSON(w, http.StatusCreated, map[string]string{"mensaje": "Caja Creada"})
}

func (c *CategoriaGasto) ListarCategoriaGasto(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	data, err := c.categoriaGastoService.ListarCategoriaGasto(ctx)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	common.ResponseJSON(w, http.StatusOK, data)
}

func (c *CategoriaGasto) ActualizarCategoriaGasto(w http.ResponseWriter, r *http.Request) {
}

func (c *CategoriaGasto) EliminarCategoriaGasto(w http.ResponseWriter, r *http.Request) {
}
