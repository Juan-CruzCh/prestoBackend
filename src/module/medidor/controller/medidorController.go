package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"prestoBackend/src/core/utils"
	"prestoBackend/src/module/medidor/dto"
	"prestoBackend/src/module/medidor/service"
	"time"

	"github.com/go-playground/validator/v10"
)

type MedidorController struct {
	service  *service.MedidorService
	Validate *validator.Validate
}

func NewMedidorController(service *service.MedidorService, Validate *validator.Validate) *MedidorController {
	return &MedidorController{
		service:  service,
		Validate: Validate,
	}

}
func (controller *MedidorController) CrearMedidor(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var body dto.MedidorDto

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	err = controller.Validate.Struct(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	resultado, err := controller.service.CrearMedidor(&body, ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resultado)
}

func (controller *MedidorController) ListarMedidorCliente(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	pagina, limite, err := utils.PaginadorHTTP(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	query := r.URL.Query()
	filter := dto.BuscadorMedidorClienteDto{
		Pagina:          pagina,
		Limite:          limite,
		Nombre:          query.Get("nombre"),
		Ci:              query.Get("ci"),
		Codigo:          query.Get("codigo"),
		ApellidoPaterno: query.Get("apellidoPaterno"),
		ApellidoMaterno: query.Get("apellidoMaterno"),
		Direccion:       query.Get("direccion"),
		NumeroMedidor:   query.Get("numeroMedidor"),
		Tarifa:          query.Get("tarifa"),
		Estado:          query.Get("estado"),
		EstadoMedidor:   query.Get("estadoMedidor"),
	}

	resultado, err := controller.service.ListarMedidores(&filter, ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)
}

func (controller *MedidorController) EliminarMedidor(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var medidor string = r.PathValue("id")

	ID, err := utils.ValidadIdMongo(medidor)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	resultado, err := controller.service.EliminarMedidor(ID, ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)
}

func (controller *MedidorController) ActualizarMedidor(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var id string = r.PathValue("id")
	ID, err := utils.ValidadIdMongo(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	var body dto.MedidorDto

	err = json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	err = controller.Validate.Struct(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	resultado, err := controller.service.ActualizarMedidor(ID, &body, ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)
}
func (controller *MedidorController) ObtenerMedidorConClientePorId(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var id string = r.PathValue("id")
	ID, err := utils.ValidadIdMongo(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	resultado, err := controller.service.ObtenerMedidorConClientePorId(ID, ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)

}
