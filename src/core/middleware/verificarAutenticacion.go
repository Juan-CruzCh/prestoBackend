package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"prestoBackend/src/core/utils"
	"prestoBackend/src/module/usuario/repository"
	"time"
)

func VerificarAutenticacion(repository repository.UsuarioRepository) func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			url := r.URL.Path
			if url == "/api/autenticacion" {
				next.ServeHTTP(w, r)
				return
			}

			cookie, err := r.Cookie("ctx")
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"mensaje": "No autenticado"})
				return
			}
			token := cookie.Value

			claims, err := utils.VerifyToken(token)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
				return
			}
			usuarioID := claims["usuario"]

			usuarioIDStr, ok := usuarioID.(string)
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"mensaje": "usuario inválido"})
				return
			}

			ID, err := utils.ValidadIdMongo(usuarioIDStr)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"mensaje": "No autenticado"})
			}

			ctxDB, cancel := context.WithTimeout(r.Context(), 10*time.Second)
			defer cancel()
			resultado, err := repository.BuscarUsuarioPorUsuarioId(ID, ctxDB)

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"mensaje": err.Error()})
			}
			dataUsuario := map[string]string{
				"id":  resultado.ID.Hex(),
				"rol": string(resultado.Rol),
			}

			ctx := context.WithValue(r.Context(), "usuario", dataUsuario)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
