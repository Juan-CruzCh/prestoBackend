package controller

import (
	"context"
	"net/http"
	"prestoBackend/src/app/common"
	"prestoBackend/src/internal/log/service"
	"time"
)

type Log struct {
	logService *service.Log
}

func NewLogController(logService *service.Log) *Log {
	return &Log{
		logService: logService,
	}
}

func (c *Log) ListarLog(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()
	resultado, err := c.logService.ListarLog(ctx)
	if err != nil {
		common.ResponseJSON(w, http.StatusBadRequest, map[string]string{"mensaje": err.Error()})
		return
	}
	common.ResponseJSON(w, http.StatusOK, resultado)
}
