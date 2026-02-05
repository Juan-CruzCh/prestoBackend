package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"prestoBackend/src/core/utils"
	"prestoBackend/src/module/pago/dto"
	"prestoBackend/src/module/pago/service"
	"time"

	"github.com/go-playground/validator/v10"
)

type PagoController struct {
	service  *service.PagoService
	Validate *validator.Validate
}

func NewPagoController(service *service.PagoService, Validate *validator.Validate) *PagoController {
	return &PagoController{
		service:  service,
		Validate: Validate,
	}
}

func (controller *PagoController) RealizarPago(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	usuarioContext := r.Context().Value("usuario")
	usuario, ok := usuarioContext.(map[string]string)
	if !ok {
		utils.ResponseJSON(w, http.StatusUnauthorized, map[string]string{"mensaje": "Usuario no encontrado en contexto"})
		return
	}
	usuarioID, err := utils.ValidadIdMongo(usuario["id"])
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	var body dto.PagoDto

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {

		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	err = controller.Validate.Struct(body)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	resultado, err := controller.service.RealizarPago(&body, usuarioID, ctx)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	utils.ResponseJSON(w, http.StatusCreated, resultado)

}

func (controller *PagoController) DetallePago(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var idPago string = r.PathValue("id")
	ID, err := utils.ValidadIdMongo(idPago)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	resultado, err := controller.service.DetallePago(ID, ctx)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	utils.ResponseJSON(w, http.StatusOK, resultado)

}

func (controller *PagoController) ListarPagos(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	pagina, limite, err := utils.PaginadorHTTP(r)

	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	query := r.URL.Query()
	var filter dto.BuscardorPagoDto = dto.BuscardorPagoDto{
		CodigoCliente:   query.Get("CodigoCliente"),
		Ci:              query.Get("ci"),
		Nombre:          query.Get("nombre"),
		ApellidoMaterno: query.Get("apellidoMaterno"),
		ApellidoPaterno: query.Get("apellidoPaterno"),
		NumeroMedidor:   query.Get("numeroMedidor"),
		FechaInicio:     query.Get("fechaInicio"),
		FechaFin:        query.Get("fechaFin"),
		Pagina:          pagina,
		Limite:          limite,
	}

	resultado, err := controller.service.ListarPagos(&filter, ctx)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	utils.ResponseJSON(w, http.StatusOK, resultado)
}
