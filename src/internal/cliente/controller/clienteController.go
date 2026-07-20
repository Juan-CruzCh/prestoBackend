package controller

import (
	"context"
	"encoding/json"
	"net/http"

	"prestoBackend/src/app/common"
	"prestoBackend/src/internal/cliente/dto"
	"prestoBackend/src/internal/cliente/service"
	"time"

	"github.com/go-playground/validator/v10"
)

type ClienteController struct {
	Service  *service.ClienteService
	Validate *validator.Validate
}

func NewClienteController(s *service.ClienteService, Validate *validator.Validate) *ClienteController {
	return &ClienteController{
		Service:  s,
		Validate: Validate,
	}
}

func (ctl *ClienteController) CrearClienteController(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var body dto.ClienteDto
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	err = ctl.Validate.Struct(body)

	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	resultado, err := ctl.Service.CrearCliente(&body, ctx)

	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	common.ResponseJSON(w, http.StatusCreated, resultado)

}

func (controller *ClienteController) ListarClientesController(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	pagina, limite, err := common.PaginadorHTTP(r)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	query := r.URL.Query()
	nombre := query.Get("nombre")
	ci := query.Get("ci")
	codigo := query.Get("codigo")
	apellidoPaterno := query.Get("apellidoPaterno")
	apellidoMaterno := query.Get("apellidoMaterno")
	celular := query.Get("celular")
	var filter dto.BucadorClienteDto = dto.BucadorClienteDto{
		Pagina:          pagina,
		Limite:          limite,
		Nombre:          nombre,
		Codigo:          codigo,
		ApellidoPaterno: apellidoPaterno,
		ApellidoMaterno: apellidoMaterno,
		Ci:              ci,
		Celular:         celular,
	}

	resultado, err := controller.Service.ListarClientes(filter, ctx)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	common.ResponseJSON(w, http.StatusOK, resultado)

}

func (controller *ClienteController) ActualizarClienteController(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var IDCliente string = r.PathValue("id")
	ID, err := common.ValidadIdMongo(IDCliente)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	var body dto.ClienteDto
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
	resultado, err := controller.Service.ActualizarCliente(&body, ID, ctx)

	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	common.ResponseJSON(w, http.StatusOK, resultado)

}

func (controller *ClienteController) EliminarClienteController(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	IDCliente := r.PathValue("id")
	ID, err := common.ValidadIdMongo(IDCliente)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	resultado, err := controller.Service.EliminarCliente(ID, ctx)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	common.ResponseJSON(w, http.StatusOK, resultado)

}
