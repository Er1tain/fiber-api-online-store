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

func FindClient(email string, password string) (string, string, bool) {
	db := connect()
	defer db.Close()

	//Пароль в форме хэш-кода
	bytePassword := []byte(password)
	md5Hash := md5.Sum(bytePassword)
	hash_password := hex.EncodeToString(md5Hash[:])

	//Поиск записи в БД
	client := Client{Email: email, Password: hash_password}
	db.First(&client)

	surname := client.Surname
	name := client.Name

	log.Println(client)
	log.Println(hash_password)

	return surname, name, client.Password == hash_password
}

func DeleteClient(client AuthData) bool {
	db := connect()
	defer db.Close()

	//Пароль в форме хэш-кода
	bytePassword := []byte(client.Password)
	md5Hash := md5.Sum(bytePassword)
	hash_password := hex.EncodeToString(md5Hash[:])

	//Поиск записи в БД
	info_client := Client{Email: client.Email, Password: hash_password}
	db.First(&info_client)

	if info_client.Name == "" || info_client.Surname == "" {
		return false
	}

	db.Delete(&info_client)
	return true

}
