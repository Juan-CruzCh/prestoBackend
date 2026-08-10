package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"prestoBackend/src/app/common"
	"prestoBackend/src/internal/caja/dto"
	"prestoBackend/src/internal/caja/service"
	"time"

	"github.com/go-playground/validator/v10"
)

type Caja struct {
	cajaService *service.Caja
	Validate    *validator.Validate
}

func NewCajaController(cajaService *service.Caja, Validate *validator.Validate) *Caja {
	return &Caja{
		cajaService: cajaService,
		Validate:    Validate,
	}
}
func (controller *Caja) CrearCaja(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var body dto.CajaDto

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}

	err = controller.Validate.Struct(&body)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}

	usuario, err := common.ObtenerUsuarioRequest(w, r)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}

	err = controller.cajaService.CrearCaja(&body, &usuario.ID, ctx)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	common.ResponseJSON(w, http.StatusCreated, map[string]string{"mensaje": "Caja Creada"})
}

func (c *Caja) ListarCaja(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	resultado, err := c.cajaService.ListarCaja(ctx)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	common.ResponseJSON(w, http.StatusOK, resultado)
}

func (c *Caja) ListarCajaPorUsuario(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	usuario, err := common.ObtenerUsuarioRequest(w, r)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	resultado, err := c.cajaService.ListarCajaPorUsuario(&usuario.ID, ctx)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	common.ResponseJSON(w, http.StatusOK, resultado)
}
