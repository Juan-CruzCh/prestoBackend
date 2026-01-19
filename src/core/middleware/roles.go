package middleware

import (
	"fmt"
	"net/http"
	"prestoBackend/src/core/enum"
)

func Roles(rol ...enum.RolE) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(rol)
			next.ServeHTTP(w, r)
		})
	}
}
