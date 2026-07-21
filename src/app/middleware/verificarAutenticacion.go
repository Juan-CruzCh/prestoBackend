package middleware

import (
	"context"
	"net/http"
	"prestoBackend/src/app/common"
	"prestoBackend/src/app/jwt"
	"prestoBackend/src/internal/usuario/repository"
	"slices"
	"time"
)

var rutasPublicas []string = []string{"/api/autenticacion"}

func VerificarAutenticacion(repository repository.UsuarioRepository) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			url := r.URL.Path
			if slices.Contains(rutasPublicas, url) {
				next.ServeHTTP(w, r)
				return
			}

			cookie, err := r.Cookie("ctx")
			if err != nil {
				common.ResponseJSON(w, http.StatusUnauthorized, (map[string]string{"mensaje": "No autenticado"}))
				return
			}
			token := cookie.Value

			claims, err := jwt.VerifyToken(token)
			if err != nil {
				common.ResponseJSON(w, http.StatusUnauthorized, (map[string]string{"mensaje": err.Error()}))
				return
			}
			usuarioID := claims["usuario"]

			usuarioIDStr, ok := usuarioID.(string)
			if !ok {
				common.ResponseJSON(w, http.StatusUnauthorized, (map[string]string{"mensaje": "usuario inválido"}))
				return
			}

			ID, err := common.ValidadIdMongo(usuarioIDStr)
			if err != nil {
				common.ResponseJSON(w, http.StatusUnauthorized, (map[string]string{"mensaje": "No autenticado"}))
			}

			ctxDB, cancel := context.WithTimeout(r.Context(), 10*time.Second)
			defer cancel()
			resultado, err := repository.BuscarUsuarioPorUsuarioId(ID, ctxDB)

			if err != nil {
				common.ResponseJSON(w, http.StatusUnauthorized, (map[string]string{"mensaje": err.Error()}))
				return
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
