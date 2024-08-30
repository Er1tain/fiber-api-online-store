package client_models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func connect() *gorm.DB {
	//Подключаемся к БД
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=ClothesShop password=userpass sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db

	//db.AutoMigrate(&Client{})
	//
	////Создаём модельку клиента
	//client := Client{Surname: "Придава", Name: "Александр", Email: "sanyapridava@mail.ru"}
	//db.Create(&client)
}

func CreateClient(client Client) {
	db := connect()
	defer db.Close()

	//Первичная инициализация таблицы клиентов
	db.AutoMigrate(&Client{})

	//Создание записи в БД
	db.Create(
		&Client{Email: client.Email, Surname: client.Surname, Name: client.Name, Password: client.Password},
	)

}
