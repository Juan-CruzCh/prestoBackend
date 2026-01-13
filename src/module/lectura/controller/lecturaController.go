package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"prestoBackend/src/core/utils"
	"prestoBackend/src/module/lectura/dto"
	"prestoBackend/src/module/lectura/service"
	"time"

	"github.com/go-playground/validator/v10"
)

type LecturaController struct {
	service *service.LecturaService
}

func NewLecturaController(service *service.LecturaService) *LecturaController {
	return &LecturaController{
		service: service,
	}
}

func (controller *LecturaController) ListarLecturas(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	validate := validator.New()

	var body dto.BuscadorLecturaDto

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	err = validate.Struct(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	resultado, err := controller.service.RepositoryLectura.ListarLectura(&body, ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)

}

func (controller *LecturaController) CrearLectura(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	validate := validator.New()

	var body dto.LecturaDto

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
		return
	}

	err = validate.Struct(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
		return
	}

	if body.LecturaAnterior > body.LecturaActual {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": "La lectura anterior no debe ser mayor a la lectura actual"})

		return

	}

	resultado, err := controller.service.CrearLectura(&body, ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resultado)
}

func (controller *LecturaController) BuscarLecturaPorNumeroMedidor(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var numeroMedidor string = r.PathValue("numeroMedidor")
	resultado, err := controller.service.BuscarLecturaPorNumeroMedidor(numeroMedidor, ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)
}

func (controller *LecturaController) BuscarLecturasPorClienteMedidor(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var cliente string = r.PathValue("cliente")
	IDCliente, err := utils.ValidadIdMongo(cliente)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
		return
	}
	resultado, err := controller.service.BuscarLecturasPorClienteMedidor(IDCliente, ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)
}

func (controller *LecturaController) DetalleLectura(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var medidor string = r.PathValue("medidor")
	var lectura string = r.PathValue("lectura")
	IDmedidor, err := utils.ValidadIdMongo(medidor)
	IDlectura, err := utils.ValidadIdMongo(lectura)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
		return
	}
	resultado, err := controller.service.DetalleLectura(IDmedidor, IDlectura, ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)
}
func (controller *LecturaController) EliminarLectura(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var IDlectura string = r.PathValue("id")
	ID, err := utils.ValidadIdMongo(IDlectura)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
		return
	}
	resultado, err := controller.service.EliminarLectura(ID, ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)
}
