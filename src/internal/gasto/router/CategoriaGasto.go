package router

import (
	"net/http"
	"prestoBackend/src/internal/gasto/controller"
)

func NewCategoriaGastoRouter(mux *http.ServeMux, controller *controller.CategoriaGasto) {
	mux.HandleFunc("POST /api/categoriaGasto/crear", controller.CrearCategoriaGasto)
	mux.HandleFunc("GET /api/categoriaGasto/listar", controller.ListarCategoriaGasto)
	//mux.HandleFunc("DELETE /api/gasto/eliminar/{id}", controller.EliminarGasto)
}
