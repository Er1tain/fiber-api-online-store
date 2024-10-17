package validation

import (
	"api/src/shared"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateJWT(token_string string) bool {
	secret_key := shared.HmacSampleSecret
	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		// Проверка метода подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("токен сгенерирован сторонним приложением")
		}
		return secret_key, nil
	})

	if err != nil {
		return false
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		return true
	}

	return false
}
