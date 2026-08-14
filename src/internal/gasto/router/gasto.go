package router

import (
	"net/http"
	"prestoBackend/src/internal/gasto/controller"
)

func NewGastoRouter(mux *http.ServeMux, controller *controller.Gasto) {
	mux.HandleFunc("POST /api/gasto/crear", controller.CrearGasto)
	mux.HandleFunc("GET /api/gasto/listar", controller.ListarGasto)
	mux.HandleFunc("DELETE /api/gasto/eliminar/{id}", controller.EliminarGasto)
}
