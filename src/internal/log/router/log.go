package router

import (
	"net/http"
	"prestoBackend/src/internal/log/controller"
)

func NewLogRouter(mux *http.ServeMux, controller *controller.Log) {
	mux.HandleFunc("GET /api/log/listar", controller.ListarLog)
}
