package client_models

import (
	"crypto/md5"
	"encoding/hex"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func connect() *gorm.DB {
	//Подключаемся к БД
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=ClothesShop password=userpass sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db
}

func CreateClient(client Client) bool {
	db := connect()
	defer db.Close()

	//Первичная инициализация таблицы клиентов
	db.AutoMigrate(&Client{})

	//Шифрование пароля
	bytePassword := []byte(client.Password)
	md5Hash := md5.Sum(bytePassword)
	hash_password := hex.EncodeToString(md5Hash[:])

	//Создание записи в БД
	result := db.Create(
		&Client{Email: client.Email, Surname: client.Surname, Name: client.Name, Password: hash_password},
	)

	return result.Error == nil

}

func FindClient(email string, password string) {
	db := connect()
	defer db.Close()

}
