package shared

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"
)

var Black_list_tokens = []string{}

func DeleteDeadTokens() {
	new_black_list_tokens := []string{}

	for _, string_token := range Black_list_tokens {
		//На момент добавления в чс время жизни токена могло не истечь
		if CheckTokenLifeTime(string_token) {
			new_black_list_tokens = append(new_black_list_tokens, string_token)
		}
	}
	Black_list_tokens = new_black_list_tokens
}

// false – мёртвый токен
func CheckTokenLifeTime(string_token string) bool {
	//Части токена отделены точкой
	parts := strings.Split(string_token, ".")

	if len(parts) != 3 {
		log.Println("Невалидный токен: " + string_token)
		return false
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		log.Println("Не удалось получить полезную нагрузку!")
	}

	var claims map[string]string
	err = json.Unmarshal(payload, &claims)
	if err != nil {
		log.Println("Не удалось сериализовать!")
	}

	//Дата создания токена в формате 2024-10-14T14:35:23.5617143+03:00
	date_created := claims["createdAt"]
	day_created, err := strconv.Atoi(date_created[8:10])
	if err != nil {
		log.Println("Не существует такого дня в месяцах!")
		return false
	}
	month_created := date_created[5:7]

	//Текущая дата
	date_current := time.Now().Format(time.RFC3339Nano)
	day_current, err := strconv.Atoi(date_current[8:10])
	month_current := date_current[5:7]

	if err != nil {
		log.Println("Не существует такого дня в месяцах!")
		return false
	}

	return !(month_created != month_current || day_current-day_created > 1)
}
