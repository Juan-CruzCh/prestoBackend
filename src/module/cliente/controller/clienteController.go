package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"prestoBackend/src/core/utils"
	"prestoBackend/src/module/cliente/dto"
	"prestoBackend/src/module/cliente/service"
	"time"

	"github.com/go-playground/validator/v10"
)

type ClienteController struct {
	Service  *service.ClienteService
	Validate *validator.Validate
}

func NewClienteController(s *service.ClienteService) *ClienteController {
	return &ClienteController{
		Service: s,
	}
}

func (ctl *ClienteController) CrearClienteController(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var body dto.ClienteDto
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
		return
	}

	err = ctl.Validate.Struct(body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
		return
	}
	resultado, err := ctl.Service.CrearCliente(&body, ctx)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resultado)
}

func (controller *ClienteController) ListarClientesController(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	pagina, limite, err := utils.PaginadorHTTP(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	query := r.URL.Query()
	nombre := query.Get("nombre")
	ci := query.Get("ci")
	codigo := query.Get("codigo")
	apellidoPaterno := query.Get("apellidoPaterno")
	apellidoMaterno := query.Get("apellidoMaterno")
	var filter dto.BucadorClienteDto = dto.BucadorClienteDto{
		Pagina:          pagina,
		Limite:          limite,
		Nombre:          nombre,
		Codigo:          codigo,
		ApellidoPaterno: apellidoPaterno,
		ApellidoMaterno: apellidoMaterno,
		Ci:              ci,
	}

	resultado, err := controller.Service.ListarClientes(filter, ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)

}

func (controller *ClienteController) ActualizarClienteController(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var IDCliente string = r.PathValue("id")
	ID, err := utils.ValidadIdMongo(IDCliente)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
		return
	}

	validate := validator.New()
	var body dto.ClienteDto
	err = json.NewDecoder(r.Body).Decode(&body)

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
	resultado, err := controller.Service.ActualizarCliente(&body, ID, ctx)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)

}

func (controller *ClienteController) EliminarClienteController(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	IDCliente := r.PathValue("id")
	ID, err := utils.ValidadIdMongo(IDCliente)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
		return
	}
	resultado, err := controller.Service.EliminarCliente(ID, ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resultado)

}
