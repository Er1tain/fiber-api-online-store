package tests_client

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"testing"
)

var (
	list_request_data = [][]byte{
		[]byte(`{
				"email": "sanyapridava@mail.ru",
				"password": "salovar"
				}`),

		[]byte(`{
				"email": "example.com@yandex.ru",
				"password": "salovar"
		}`),
	}
)

// Тест аутентификации
func TestAuthHandler(t *testing.T) {
	for _, request_data := range list_request_data {
		r := bytes.NewReader(request_data)

		response, err := http.Post("http://localhost:8000/api/client/auth", "application/json", r)
		if err != nil {
			log.Println("Не удалось отправить тестовый запрос на аутентификацию(")
		}
		defer response.Body.Close()

		body, _ := io.ReadAll(response.Body)

		actual := string(body)
		log.Println(actual)

		//Сравним полученный и ожидаемый результаты
		succes_1 := "Неверный логин и/или пароль!"
		if actual != succes_1 {
			t.Errorf("Result was incorrect, got: %s, want: %s.", actual, succes_1)
		}

	}
}
