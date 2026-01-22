package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"prestoBackend/src/core/utils"
	"prestoBackend/src/module/tarifa/dto"
	"prestoBackend/src/module/tarifa/service"
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
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)
}

func (controller *TarifaController) ListarTarifas(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)

	defer cancel()
	resultado, err := controller.tarifaService.ListarTarifas(ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)

}
func (controller *TarifaController) CrearTarifa(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var body dto.TarifaDto
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
		return
	}
	err = controller.Validate.Struct(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
		return
	}

	resultado, err := controller.tarifaService.CrearTarifa(&body, ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resultado)
}

func (controller *TarifaController) EliminarTarifa(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var tarifa string = r.PathValue("id")
	taridaId, err := utils.ValidadIdMongo(tarifa)
	resultado, err := controller.tarifaService.EliminarTarifa(taridaId, ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)
}

func (controller *TarifaController) EliminarRango(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var rango string = r.PathValue("id")
	rangoId, err := utils.ValidadIdMongo(rango)
	resultado, err := controller.tarifaService.EliminarRango(rangoId, ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)

}
