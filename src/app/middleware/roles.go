package middleware

import (
	"encoding/json"
	"net/http"
	"prestoBackend/src/app/enum"
)

func Roles(rol ...enum.RolE) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			contextUsuario := r.Context().Value("usuario")
			usuario, ok := contextUsuario.(map[string]string)
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{"mensaje": "Usuario mal formado"})
			}
			for _, v := range rol {
				if string(v) == usuario["rol"] {
					next.ServeHTTP(w, r)
					return
				}
			}
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]string{"mensaje": "no autorizado"})
		})
	}
}
