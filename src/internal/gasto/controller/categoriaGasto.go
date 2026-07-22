package controller

import (
	"net/http"
	"prestoBackend/src/internal/gasto/service"
)

type CategoriaGasto struct {
	categoriaGastoService *service.CategoriaGasto
}

func NewCategoriaGastoController(categoriaGastoService *service.CategoriaGasto) *CategoriaGasto {
	return &CategoriaGasto{
		categoriaGastoService: categoriaGastoService,
	}
}
func (c *CategoriaGasto) CrearCategoriaGasto(w http.ResponseWriter, r *http.Request) {
}

func (c *CategoriaGasto) ListarCategoriaGasto(w http.ResponseWriter, r *http.Request) {
}

func (c *CategoriaGasto) ActualizarCategoriaGasto(w http.ResponseWriter, r *http.Request) {
}

func (c *CategoriaGasto) EliminarCategoriaGasto(w http.ResponseWriter, r *http.Request) {
}
