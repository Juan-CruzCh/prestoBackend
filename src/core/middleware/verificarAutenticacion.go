package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"prestoBackend/src/core/utils"
)

func VerificarAutenticacion(next http.Handler) http.Handler {
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
		ctx := context.WithValue(r.Context(), "usuario", usuarioID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
