package common

import "github.com/go-playground/validator/v10"

func ErrorJson(err error) []map[string]string {
	errores := make([]map[string]string, 0)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {

		for _, e := range validationErrors {
			errorMsg := e.Tag()
			if e.Param() != "" {
				errorMsg += "=" + e.Param()
			}

			errores = append(errores, map[string]string{
				"field": e.Field(),
				"error": errorMsg,
			})
		}

		return errores
	}
	return errores
}
