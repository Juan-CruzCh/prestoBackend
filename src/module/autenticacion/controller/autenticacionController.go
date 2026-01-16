package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"prestoBackend/src/module/autenticacion/dto"
	"prestoBackend/src/module/autenticacion/service"
	"time"

	"github.com/go-playground/validator/v10"
)

type AutenticacionController struct {
	service *service.AutenticacionService
}

func NewAutenticacionController(service *service.AutenticacionService) *AutenticacionController {
	return &AutenticacionController{
		service: service,
	}
}

func (controller *AutenticacionController) Autenticacion(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	validate := validator.New()
	var body dto.AutenticacionDto
	fmt.Println(body)
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
		return
	}

	err = validate.Struct(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
		return
	}

	token, err := controller.service.Autenticacion(&body, ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "ctx",
		Value: token,
		Path:  "/",
		//Domain:   "http://localhost:4200",
		MaxAge:   4 * 60 * 60,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})

}
