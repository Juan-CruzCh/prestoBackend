package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var claveSecreta []byte = []byte("hola soy la clave sercrea")

func GenraraToken(usuario bson.ObjectID) (string, error) {
	claims := jwt.MapClaims{
		"usuario": usuario,
		"exp":     time.Now().Add(time.Hour * 4).Unix(),
		"iat":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(claveSecreta)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return claveSecreta, nil
	})

	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, errors.New("jwt invalido")
	}
	return claims, nil

}
