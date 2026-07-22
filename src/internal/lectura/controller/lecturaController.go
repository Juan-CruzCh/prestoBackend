package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"prestoBackend/src/app/common"
	"prestoBackend/src/internal/lectura/dto"
	"prestoBackend/src/internal/lectura/service"
	"time"

	"github.com/go-playground/validator/v10"
)

type LecturaController struct {
	service  *service.LecturaService
	Validate *validator.Validate
}

func NewLecturaController(service *service.LecturaService, Validate *validator.Validate) *LecturaController {
	return &LecturaController{
		service:  service,
		Validate: Validate,
	}
}

func (controller *LecturaController) ListarLecturas(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var body dto.BuscadorLecturaDto

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

	resultado, err := controller.service.ListarLectura(&body, ctx)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	common.ResponseJSON(w, http.StatusOK, resultado)

}

func (controller *LecturaController) CrearLectura(w http.ResponseWriter, r *http.Request) {
	usuarioContext := r.Context().Value("usuario")
	usuario, ok := usuarioContext.(map[string]string)
	if !ok {
		common.ResponseJSON(w, http.StatusUnauthorized, map[string]string{"mensaje": "Usuario no encontrado en contexto"})

		return
	}

	idUsuario, err := common.ValidadIdMongo(usuario["id"])
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var body dto.LecturaDto

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}

	err = controller.Validate.Struct(&body)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}

	resultado, err := controller.service.CrearLectura(&body, idUsuario, ctx)

	if err != nil {
		if err.Error() == "La lectura ya se encuentra registrada" {
			common.ResponseJSON(w, http.StatusConflict, map[string]string{"mensaje": err.Error()})
			return
		} else {
			common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
			return
		}

	}

	common.ResponseJSON(w, http.StatusCreated, resultado)
}

func (controller *LecturaController) BuscarLecturaPorNumeroMedidor(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var numeroMedidor string = r.PathValue("numeroMedidor")
	resultado, err := controller.service.BuscarLecturaPorNumeroMedidor(numeroMedidor, ctx)
	if err != nil {
		if err.Error() == "Numero de medidor no encontrado" {
			common.ResponseJSON(w, http.StatusNotFound, map[string]string{"mensaje": err.Error()})
			return
		} else {
			common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
			return

		}

	}
	common.ResponseJSON(w, http.StatusOK, resultado)

}

func (controller *LecturaController) BuscarLecturasPorClienteMedidor(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var cliente string = r.PathValue("cliente")
	IDCliente, err := common.ValidadIdMongo(cliente)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	resultado, err := controller.service.BuscarLecturasPorClienteMedidor(IDCliente, ctx)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	common.ResponseJSON(w, http.StatusOK, resultado)
}

func (controller *LecturaController) DetalleLectura(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var medidor string = r.PathValue("medidor")
	var lectura string = r.PathValue("lectura")
	IDmedidor, err := common.ValidadIdMongo(medidor)
	IDlectura, err := common.ValidadIdMongo(lectura)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	resultado, err := controller.service.DetalleLectura(IDmedidor, IDlectura, ctx)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	common.ResponseJSON(w, http.StatusOK, resultado)
}
func (controller *LecturaController) EliminarLectura(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	var IDlectura string = r.PathValue("id")
	ID, err := common.ValidadIdMongo(IDlectura)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	resultado, err := controller.service.EliminarLectura(ID, ctx)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	common.ResponseJSON(w, http.StatusOK, resultado)
}
