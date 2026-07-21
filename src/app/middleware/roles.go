package middleware

import (
	"net/http"
	"prestoBackend/src/app/common"
	"prestoBackend/src/app/enum"
)

func Roles(rol ...enum.RolE) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			contextUsuario := r.Context().Value("usuario")
			usuario, ok := contextUsuario.(map[string]string)
			if !ok {
				common.ResponseJSON(w, http.StatusUnauthorized, (map[string]string{"mensaje": "Usuario mal formado"}))
				return
			}
			for _, v := range rol {
				if string(v) == usuario["rol"] {
					next.ServeHTTP(w, r)
					return
				}
			}
			common.ResponseJSON(w, http.StatusForbidden, (map[string]string{"mensaje": "no autorizado"}))
		})
	}
}
