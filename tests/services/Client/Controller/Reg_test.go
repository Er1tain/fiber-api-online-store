package tests_client

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"testing"
)

// Тест регистрации
func TestRegHandler(t *testing.T) {
	request_data := []byte(`{
			"email": "saruma2n@gmail.com",
			"surname": "Tar6on",
			"name": "Jame3s",
			"password": "mega1don1"
		}`)

	r := bytes.NewReader(request_data)

	response, err := http.Post("http://localhost:8000/api/client/reg", "application/json", r)
	if err != nil {
		log.Println("Не удалось отправить тестовый запрос на регистрацию(")
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	actual := string(body)
	log.Println(actual)

	//Сравним полученный и ожидаемый результаты
	expected := "Данный email уже используется"
	if expected != actual {
		t.Errorf("Result was incorrect, got: %s, want: %s.", actual, expected)
	}

}
