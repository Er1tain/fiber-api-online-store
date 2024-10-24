package profile_middleware

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"strings"
)

// Получить email из jwt
func GetUserEmail(token_string string) (string, error) {
	//Части токена отделены точкой
	parts := strings.Split(token_string, ".")

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		msg := "Не удалось получить полезную нагрузку!"
		log.Println(msg)
		return msg, err
	}

	var claims map[string]string
	err = json.Unmarshal(payload, &claims)
	if err != nil {
		msg := "Не удалось сериализовать!"
		log.Println(msg)
		return msg, err
	}

	return claims["email"], nil
}

// Получить ФИ
type surname = string
type name = string

func GetUserSurnameName(token_string string) (surname, name, error) {
	//Части токена отделены точкой
	parts := strings.Split(token_string, ".")

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		msg := "Не удалось получить полезную нагрузку!"
		log.Println(msg)
		return msg, "", err
	}

	var claims map[string]string
	err = json.Unmarshal(payload, &claims)
	if err != nil {
		msg := "Не удалось сериализовать!"
		log.Println(msg)
		return msg, "", err
	}

	surname := claims["surname"]
	name := claims["name"]

	return surname, name, nil
}
