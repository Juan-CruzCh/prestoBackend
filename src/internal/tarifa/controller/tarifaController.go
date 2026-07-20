package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"prestoBackend/src/app/common"
	"prestoBackend/src/internal/tarifa/dto"
	"prestoBackend/src/internal/tarifa/service"
	"time"

	"github.com/go-playground/validator/v10"
)

type TarifaController struct {
	tarifaService *service.TarifaService
	Validate      *validator.Validate
}

func NewTarifaController(tarifaService *service.TarifaService, Validate *validator.Validate) *TarifaController {
	return &TarifaController{
		tarifaService: tarifaService,
		Validate:      Validate,
	}
}

func (controller *TarifaController) ListarTarifasConRagos(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)

	defer cancel()
	resultado, err := controller.tarifaService.ListarTarifasConRagos(ctx)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	common.ResponseJSON(w, http.StatusOK, resultado)
}

func (controller *TarifaController) ListarTarifas(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)

	defer cancel()
	resultado, err := controller.tarifaService.ListarTarifas(ctx)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	common.ResponseJSON(w, http.StatusOK, resultado)

}
func (controller *TarifaController) CrearTarifa(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var body dto.TarifaDto
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	err = controller.Validate.Struct(&body)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	resultado, err := controller.tarifaService.CrearTarifa(&body, ctx)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	common.ResponseJSON(w, http.StatusCreated, resultado)
}

func (controller *TarifaController) EditarTarifa(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var tarifaId string = r.PathValue("id")
	tarifa, err := common.ValidadIdMongo(tarifaId)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	var body dto.TarifaDto
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	err = controller.Validate.Struct(&body)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	resultado, err := controller.tarifaService.EditarTarifa(&body, tarifa, ctx)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	common.ResponseJSON(w, http.StatusOK, resultado)
}

func (controller *TarifaController) EliminarTarifa(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var tarifa string = r.PathValue("id")
	taridaId, err := common.ValidadIdMongo(tarifa)
	resultado, err := controller.tarifaService.EliminarTarifa(taridaId, ctx)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	common.ResponseJSON(w, http.StatusOK, resultado)
}

func (controller *TarifaController) ObtenerTarifasRangosId(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var tarifaId string = r.PathValue("id")
	tarifa, err := common.ValidadIdMongo(tarifaId)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	resultado, err := controller.tarifaService.ObtenerTarifasRangosId(tarifa, ctx)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	common.ResponseJSON(w, http.StatusOK, resultado)

}
