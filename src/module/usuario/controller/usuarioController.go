package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"prestoBackend/src/core/utils"
	coreUtils "prestoBackend/src/core/utils"
	"prestoBackend/src/module/usuario/dto"
	"prestoBackend/src/module/usuario/service"
	"time"

	"github.com/go-playground/validator/v10"
)

type UsuarioController struct {
	service  *service.UsuarioService
	Validate *validator.Validate
}

func NewUsuarioController(service *service.UsuarioService, Validate *validator.Validate) *UsuarioController {
	return &UsuarioController{
		service:  service,
		Validate: Validate,
	}
}

func (controller *UsuarioController) CrearUsuarios(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	validate := validator.New()
	var body dto.UsuarioDto

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})

		return
	}
	err = validate.Struct(body)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})

		return
	}

	resultado, err := controller.service.CrearUsuario(&body, ctx)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})

		return
	}
	utils.ResponseJSON(w, http.StatusCreated, resultado)

}

func (controller *UsuarioController) ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	resultado, err := controller.service.ListarUsuarios(ctx)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})

		return
	}
	utils.ResponseJSON(w, http.StatusOK, resultado)
}

func (controller *UsuarioController) Eliminar(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var id string = r.PathValue("id")
	ID, err := coreUtils.ValidadIdMongo(id)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})

		return
	}
	resultado, err := controller.service.Eliminar(ID, ctx)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})

		return
	}
	utils.ResponseJSON(w, http.StatusOK, resultado)

}
func (controller *UsuarioController) ActualizarUsuarios(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	var id string = r.PathValue("id")
	ID, err := coreUtils.ValidadIdMongo(id)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})

		return
	}
	validate := validator.New()
	var body dto.UsuarioDto

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})

		return
	}
	err = validate.Struct(body)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})

		return
	}
	resultado, err := controller.service.ActualizarUsuario(ID, &body, ctx)
	if err != nil {
		utils.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})

		return
	}
	utils.ResponseJSON(w, http.StatusOK, resultado)

}
