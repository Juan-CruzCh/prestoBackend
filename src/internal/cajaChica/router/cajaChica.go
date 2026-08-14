package router

import (
	"net/http"
	"prestoBackend/src/internal/cajaChica/controller"
)

func NewCajaChicaRouter(mux *http.ServeMux, controller *controller.CajaChica) {
	mux.HandleFunc("POS /api/cajaChica/crear", controller.CrearCajaChica)
	mux.HandleFunc("POS /api/cajaChica/listar", controller.ListarCajaChica)
}
