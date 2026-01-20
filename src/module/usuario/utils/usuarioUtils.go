package utils

import "regexp"

func ValidarContrasena(password string) bool {
	pattern := `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*])[A-Za-z\d!@#$%^&*]{8,}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(password)
}
